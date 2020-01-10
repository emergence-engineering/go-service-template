SERVICE_NAME := "example-service"
GIT_COMMIT_ID := $(shell git describe --tags --always)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
LINKFLAGS :=-s \
	-X github.com/emergence-engineering/go-service-template/internal/service.GIT_COMMIT_ID=$(GIT_COMMIT_ID) \
	-X github.com/emergence-engineering/go-service-template/internal/service.GIT_BRANCH=$(GIT_BRANCH) \
	-X github.com/emergence-engineering/go-service-template/internal/service.SERVICE_NAME=$(SERVICE_NAME) \
	-X github.com/emergence-engineering/go-service-template/internal/service.BUILD_TIMESTAMP=$(shell date -u '+%Y-%m-%dT%H:%M:%S%z')

build:
	cd cmd && CGO_ENABLED=0 go build -ldflags "$(LINKFLAGS)" -o ../$(SERVICE_NAME) -a .

test:
	go vet .
	go test -race
