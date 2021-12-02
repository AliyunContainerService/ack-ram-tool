
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
VERSION ?= $(shell git describe --tags --abbrev=0)
CGO_ENABLED ?= 0
LDFLAGS := -extldflags "-static"
LDFLAGS += -X github.com/AliyunContainerService/ack-ram-tool/pkg/version.Version=$(VERSION)
LDFLAGS += -X github.com/AliyunContainerService/ack-ram-tool/pkg/version.GitCommit=$(GIT_COMMIT)

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) go build -ldflags "$(LDFLAGS)" -a -o ack-ram-tool \
	cmd/ack-ram-tool/main.go
