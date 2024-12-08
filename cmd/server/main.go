package main

import (
	"flag"
	"kv-store/internal/app"
	"kv-store/internal/config"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Parse command-line flags
	populate := flag.Bool("populate", false, "Populate the store with test data")
	flag.Parse()

	if *populate {
		populateStore(cfg)
		return
	}

	app.Run(cfg)
}
