
# Standard build and test targets for the project
.PHONY: build
build:
	go build -o go-superstruct-port ./...

.PHONY: test
test:
	go test -race ./...

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: bench
bench:
	go test -bench=. ./...

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
