.PHONY: build
build:
		go build -v ./cmd/yukat

.PHONY: migrations
migrations-up:
		psql sardor -h localhost -p 5432 -d yukatdb -f ./migrations/01_create_tables_up.sql

migrations-down:
		psql sardor -h localhost -p 5432 -d yukatdb -f ./migrations/01_create_tables_down.sql

.DEFAULT_GOAL:=build