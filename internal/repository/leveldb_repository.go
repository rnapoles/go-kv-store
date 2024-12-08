package repository

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBRepository struct {
	db *leveldb.DB
}

func NewLevelDBRepository(path string) (*LevelDBRepository, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBRepository{db: db}, nil
}

func (r *LevelDBRepository) Set(key, value string) error {
	return r.db.Put([]byte(key), []byte(value), nil)
}

func (r *LevelDBRepository) Get(key string) (string, error) {
	data, err := r.db.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *LevelDBRepository) Delete(key string) error {
	return r.db.Delete([]byte(key), nil)
}

func (r *LevelDBRepository) List() (map[string]string, error) {
	result := make(map[string]string)

	iter := r.db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		result[key] = value
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	return result, nil
}
