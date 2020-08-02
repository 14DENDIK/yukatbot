.PHONY: build
build:
		go build -v ./cmd/yukat

.PHONY: migrations
migrations-down:
		psql sardor -h localhost -p 5432 -d yukatdb -f ./migrations/01_create_tables_down.sql
		
migrations-up:
		psql sardor -h localhost -p 5432 -d yukatdb -f ./migrations/01_create_tables_up.sql

.PHONY: generate
generate:
		go generate ./cmd/yukat

.DEFAULT_GOAL:=build