package main

import (
	"kv-store/internal/config"
	"kv-store/internal/infrastructure"
	"log"
)

func populateStore(cfg *config.Config) {

	// Choose the repository based on the database type
	repo, err := infrastructure.InitRepository(cfg.DBType, cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Add test data
	testData := map[string]string{
		"username": "john_doe",
		"email":    "john@example.com",
		"role":     "admin",
		"active":   "true",
	}

	for key, value := range testData {
		if err := repo.Set(key, value); err != nil {
			log.Printf("Failed to set key '%s': %v", key, err)
		}
	}

	log.Println("Store populated with test data successfully.")
}
