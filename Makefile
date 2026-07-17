build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
.PHONY: build

lint:
	golangci-lint run
.PHONY: lint

lint-fix:
	golangci-lint run --fix
.PHONY: lint-fix
