package app

import (
	"fmt"
	"kv-store/internal/config"
	"kv-store/internal/delivery/controller"
	"kv-store/internal/infrastructure"
	"kv-store/internal/usecase"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {

	repo, err := infrastructure.InitRepository(cfg.DBType, cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	log.Printf("Using database: %s\n", cfg.DBType)

	service := usecase.NewKeyValueService(repo)
	handler := controller.NewKeyValueHandler(service)

	http.HandleFunc("/set", handler.Set)
	http.HandleFunc("/get", handler.Get)
	http.HandleFunc("/delete", handler.Delete)
	http.HandleFunc("/list", handler.List)

	log.Printf("Server running on port: %s\n", cfg.HTTPPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTPPort), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
