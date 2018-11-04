package store

import (
	"database/sql"
	"fmt"
)

func needsMigration(store *DbStore) (bool, error) {
	row := store.db.QueryRow("SELECT 1 FROM information_schema.tables WHERE table_name = 'schemaversionw'")
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
		fmt.Println(version)
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
	return nil
}
