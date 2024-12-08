package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(path string) (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS keyvalue (key TEXT PRIMARY KEY, value TEXT)")
	if err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) Set(key, value string) error {
	_, err := r.db.Exec("INSERT OR REPLACE INTO keyvalue (key, value) VALUES (?, ?)", key, value)
	return err
}

func (r *SQLiteRepository) Get(key string) (string, error) {
	var value string
	err := r.db.QueryRow("SELECT value FROM keyvalue WHERE key = ?", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (r *SQLiteRepository) Delete(key string) error {
	_, err := r.db.Exec("DELETE FROM keyvalue WHERE key = ?", key)
	return err
}

func (r *SQLiteRepository) List() (map[string]string, error) {
	result := make(map[string]string)

	rows, err := r.db.Query("SELECT key, value FROM keyvalue")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		result[key] = value
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
