# 当前目录
CUR_DIR=$(shell pwd)
OUT_DIR=$(CUR_DIR)/bin

# 命令
GO_BUILD = CGO_ENABLED=0 go build -trimpath

TOOL_VERSION	?= $(shell git describe --long --tags --dirty --always)
TOOL_VERSION	?= unkonwn
BUILD_TIME      ?= $(shell date "+%F_%T_%Z")
COMMIT_SHA1     ?= $(shell git show -s --format=%h)
COMMIT_LOG      ?= $(shell git show -s --format=%s)
COMMIT_AUTHOR	?= $(shell git show -s --format=%an)
COMMIT_DATE		?= $(shell git show -s --format=%ad)
VERSION_MSG		?= $(COMMIT_AUTHOR)|$(COMMIT_DATE)|${COMMIT_LOG}
VERSION_PACKAGE	?= github.com/zhangga/luban/pkg/version

.PHONY: lint
# run all lint
lint:
	golangci-lint run -c .golangci.yml ./...

.PHONY: test
# run all test
test:
	go test -race ./...

VERSION_BUILD_LDFLAGS= \
-X "${VERSION_PACKAGE}.Version=${TOOL_VERSION}" \
-X "${VERSION_PACKAGE}.Message=${VERSION_MSG}" \
-X "${VERSION_PACKAGE}.BuildTime=${BUILD_TIME}" \
-X "${VERSION_PACKAGE}.CommitHash=${COMMIT_SHA1}"
.PHONY: build
# build
build:
	$(GO_BUILD) \
	-ldflags '$(VERSION_BUILD_LDFLAGS)' \
	-o $(OUT_DIR)/ \
	./cmd/luban


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help