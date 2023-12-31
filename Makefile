.DEFAULT_GOAL := build
SHELL := /bin/bash
.PHONY: fmt build

install:
	cd service && go get
	pnpm install

fmt:
	cd service && go fmt ./...

tidy:
	cd service && go mod tidy

build:
	cd service && go build 
	pnpm run build

start-backend:
	cd service && air


start-frontend:
	pnpm run dev

migrate: export GOOSE_DRIVER=postgres
migrate: 
	cd service && source .env && \
	GOOSE_DBSTRING=postgres://$$POSTGRES_USER@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB goose -dir db/migrations up

undo-last-migration: export GOOSE_DRIVER=postgres
undo-last-migration:
	cd service && source .env && \
	GOOSE_DBSTRING=postgres://$$POSTGRES_USER@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB goose down

sql:
	cd service && sqlc generate

preview: build start-backend
