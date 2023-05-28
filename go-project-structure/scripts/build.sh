#!/bin/bash

# Clear any existing build artifacts
rm -rf build/

# Build the API binary
go build -o build/main ./cmd

# Copy necessary configuration files or assets, if any
# cp -R config/ build/
# Add more files to copy if needed

# Optionally, run any necessary database migrations
# go run migrations/migrate.go

# Provide executable permissions to the binary
chmod +x build/main

echo "Build completed successfully."
