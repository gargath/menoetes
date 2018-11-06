package store

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	mig "github.com/gargath/menoetes/pkg/store/migrations"
)

func needsMigration(store *DbStore) (bool, error) {
	row := store.db.QueryRow("SELECT 1 FROM information_schema.tables WHERE table_name = 'schemaversion'")
	switch err := row.Scan(); err {
	case sql.ErrNoRows:
		err2 := createInitialSchema(store)
		if err2 != nil {
			return false, fmt.Errorf("schema not present but: %s", err2)
		}
	}
	var version int64
	row = store.db.QueryRow("SELECT version FROM schemaversion")
	switch err := row.Scan(&version); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		if version < mig.Latest {
			log.WithFields(log.Fields{
				"current_schema": version,
				"latest_schema":  mig.Latest,
			}).Info("schema migration required")
			migrateToLatest(store)
		}
	default:
		return false, fmt.Errorf("failed to query schema version: %s", err)
	}
	return false, nil
}

func createInitialSchema(store *DbStore) error {
	_, err := store.db.Exec("CREATE TABLE schemaversion ( version integer )")
	if err != nil {
		return fmt.Errorf("failed to create initial schema: %s", err)
	}
	_, err = store.db.Exec("INSERT INTO schemaversion (version) VALUES (0)")
	if err != nil {
		return fmt.Errorf("failed to insert initial schema version: %s", err)
	}
	return nil
}

func migrateToLatest(store *DbStore) error {
	var version int64
	row := store.db.QueryRow("SELECT version FROM schemaversion")
	err := row.Scan(&version)
	if err != nil {
		return fmt.Errorf("failed to get schema version: %s", err)
	}
	for version != mig.Latest {
		stmts := mig.MigrationsFor(version)
		for _, st := range stmts {
			log.WithFields(log.Fields{
				"current_statmenet": st,
				"total_statements":  len(stmts),
			}).Info("migration in progress")
			_, err = store.db.Exec(st)
			if err != nil {
				fmt.Errorf("schema migration failed! Schema may be corrupted: %s", err)
			}
		}
		row := store.db.QueryRow("SELECT version FROM schemaversion")
		err := row.Scan(&version)
		if err != nil {
			return fmt.Errorf("failed to get schema version: %s", err)
		}
		log.Info(fmt.Errorf("Schema migration step complete. Schema version now: %d", version))
	}
	log.Info("Schema migration finished")
	return nil
}
