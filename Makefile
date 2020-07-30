.PHONY: build
build:
		go build -v ./cmd/yukat

.DEFAULT_GOAL:=build