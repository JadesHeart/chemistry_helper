package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config *DataBaseConfig
	db     *sql.DB
}

func NewStore(config *DataBaseConfig) *Store {
	return &Store{
		config: config,
	}
}

func (store *Store) Open() error {
	db, err := sql.Open("postgres", store.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	store.db = db
	return nil
}

func (store *Store) Close() {
	store.db.Close()
}

func (store *Store) GetterDB() *sql.DB {
	return store.db
}
