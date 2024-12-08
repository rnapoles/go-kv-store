package usecase

import (
	"errors"
	"kv-store/internal/domain"
)

type KeyValueService struct {
	repo domain.KeyValueRepository
}

func NewKeyValueService(repo domain.KeyValueRepository) *KeyValueService {
	return &KeyValueService{repo: repo}
}

func (s *KeyValueService) Set(key, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	return s.repo.Set(key, value)
}

func (s *KeyValueService) Get(key string) (string, error) {
	return s.repo.Get(key)
}

func (s *KeyValueService) Delete(key string) error {
	return s.repo.Delete(key)
}

func (s *KeyValueService) List() (map[string]string, error) {
	return s.repo.List()
}
