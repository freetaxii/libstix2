#!/bin/bash

# Script to run all tests except SQLite datastore tests

set -e

echo "Running tests (excluding SQLite datastore tests)..."

# Run tests excluding the sqlite3 datastore directory
# We use 'go test ./...' to test all packages and then exclude the sqlite3 package
go test $(go list ./... | grep -v github.com/freetaxii/libstix2/datastore/sqlite3)

echo "All non-SQLite tests completed successfully!"