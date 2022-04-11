PWD = $(shell pwd)
NAME = bmstu-fv-parser

.PHONY: run
run:
	go run $(PWD)/cmd/$(NAME)/

.PHONY: build
build:
	go build -o bin/$(NAME) $(PWD)/cmd/$(NAME)

# Запустить тесты
.PHONY: test
test:
	go test $(PWD)/... -parallel=10 -coverprofile=cover.out

.PHONY: local
local:
	cp .dist.env .env