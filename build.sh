#!/bin/bash

set -e

echo "Cleaning prev build artifacts"

rm -rf bin
mkdir -p bin

go build -o bin/server.exe ./src/server
go build -o bin/client.exe ./src/client

echo "Build complete"