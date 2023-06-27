.DEFAULT_GOAL := build
SHELL := /bin/bash
.PHONY: fmt build

fmt:
	cd service && go fmt ./...

tidy:
	cd service && go mod tidy

build:
	cd service && go build 

start-backend:
	cd service && ./service

start-frontend:
	cd frontend && pnpm run dev

migrate: export GOOSE_DRIVER=postgres
migrate: 
	cd service && source .env && \
	GOOSE_DBSTRING=postgres://$$POSTGRES_USER@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB goose -dir db/migrations up

sql:
	cd service && sqlc generate