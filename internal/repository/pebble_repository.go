package repository

import (
	"github.com/cockroachdb/pebble"
)

type PebbleRepository struct {
	db *pebble.DB
}

func NewPebbleRepository(path string) (*PebbleRepository, error) {
	db, err := pebble.Open(path, nil)
	if err != nil {
		return nil, err
	}
	return &PebbleRepository{db: db}, nil
}

func (r *PebbleRepository) Set(key, value string) error {
	return r.db.Set([]byte(key), []byte(value), pebble.Sync)
}

func (r *PebbleRepository) Get(key string) (string, error) {
	data, closer, err := r.db.Get([]byte(key))
	if err != nil {
		return "", err
	}
	defer closer.Close()
	return string(data), nil
}

func (r *PebbleRepository) Delete(key string) error {
	return r.db.Delete([]byte(key), pebble.Sync)
}

func (r *PebbleRepository) List() (map[string]string, error) {
	result := make(map[string]string)

	iter, err := r.db.NewIter(nil)
	if err != nil {
		return nil, err
	}

	defer iter.Close()

	for iter.First(); iter.Valid(); iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		result[key] = value
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	return result, nil
}
