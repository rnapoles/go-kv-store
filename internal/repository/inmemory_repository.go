package repository

import (
	"errors"
	"sync"
)

type InMemoryRepository struct {
	store map[string]interface{}
	mu    sync.RWMutex // Mutex to ensure thread safety
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		store: make(map[string]interface{}),
	}
}

// Set stores a key-value pair in memory
func (r *InMemoryRepository) Set(key string, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[key] = value
	return nil
}

// Get retrieves a value by key
func (r *InMemoryRepository) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	value, exists := r.store[key]
	if !exists {
		return "", errors.New("key not found")
	}

	strValue, ok := value.(string)
	if !ok {
		return "", errors.New("value is not a string")
	}

	return strValue, nil
}

// Delete removes a key-value pair by key
func (r *InMemoryRepository) Delete(key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.store[key]; !exists {
		return errors.New("key not found")
	}

	delete(r.store, key)
	return nil
}

// List returns all key-value pairs in memory
func (r *InMemoryRepository) List() (map[string]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Create a copy to ensure safety and avoid external modifications
	result := make(map[string]string)
	for key, value := range r.store {
		strValue, ok := value.(string)
		if !ok {
			continue // Ignore non-string values for safety
		}
		result[key] = strValue
	}

	return result, nil
}
