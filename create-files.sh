#!/bin/bash

# Create directories
echo "Creating directories..."
mkdir -p cmd/server \
         internal/config \
         internal/domain \
         internal/usecase \
         internal/repository \
         internal/delivery/http \
         internal/infrastructure \
         internal/app

# Create files
echo "Creating Go files..."

# Entry point
touch cmd/server/main.go

# Config
touch internal/config/config.go

# Domain
touch internal/domain/keyvalue.go

# Usecase
touch internal/usecase/keyvalue_service.go

# Repositories
touch internal/repository/leveldb_repository.go
touch internal/repository/pebble_repository.go
touch internal/repository/sqlite_repository.go

# Delivery (HTTP Handlers)
touch internal/delivery/http/keyvalue_handler.go

# Infrastructure
touch internal/infrastructure/database.go

# App
touch internal/app/app.go

echo "All files and directories created successfully!"
