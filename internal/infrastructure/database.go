package infrastructure

import (
	"fmt"
	"kv-store/internal/domain"
	"kv-store/internal/repository"
)

func InitRepository(dbType, dbPath string) (domain.KeyValueRepository, error) {
	switch dbType {
	case "leveldb":
		return repository.NewLevelDBRepository(dbPath)
	case "pebble":
		return repository.NewPebbleRepository(dbPath)
	case "sqlite":
		return repository.NewSQLiteRepository(dbPath)
	case "memory":
		return repository.NewInMemoryRepository(), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
