#!/usr/bin/env bash

set -euo pipefail

GOCACHE=off go test -v -race
go fmt
go vet
golint
cd pkg
GOCACHE=off go test -v -race ./...
go fmt ./...
go vet ./...
golint ./...