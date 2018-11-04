package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

type DbStore struct {
	db *sql.DB
}

func (s *DbStore) GetLatestVersion(namespace string, name string, provider string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func (s *DbStore) GetDownloadURL(namespace string, name string, provider string, version string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func (s *DbStore) GetModulesList(namespace string, name string, provider string) ([]Module, error) {
	return []Module{}, fmt.Errorf("Not implemented")
}

func (s *DbStore) GetModuleVersions(namespace string, name string, provider string) ([]Module, error) {
	return []Module{}, fmt.Errorf("Not implemented")
}

func (s *DbStore) GetModuleDetails(namespace string, name string, provider string, version string) (Module, error) {
	return Module{}, fmt.Errorf("Not implemented")
}

func (s *DbStore) ValidateAccessToken(token string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func NewDbStore() (Store, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.username"), viper.GetString("database.password"), viper.GetString("database.dbname"))
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}
	store := &DbStore{db: db}
	outdated, err := needsMigration(store)
	if err != nil {
		return nil, fmt.Errorf("failed to verify if schema migration is required: %s", err)
	}
	if outdated {
		err := migrateToLatest(store)
		if err != nil {
			return nil, err
		}
	}
	return &DbStore{db: db}, nil
}
