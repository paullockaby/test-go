# stop on error, no built in rules, run silently
MAKEFLAGS="-S -s -r"

# get tag and commit information
IMAGE_COMMIT := $(shell git log -1 | head -n 1 | cut -d" " -f2)
IMAGE_TAG := $(shell git tag --contains ${IMAGE_COMMIT})

# set the version from the tag and commit details
IMAGE_VERSION := $(or $(IMAGE_TAG),$(IMAGE_COMMIT))
ifneq ($(shell git status --porcelain),)
    IMAGE_VERSION := $(IMAGE_VERSION)-dirty
endif

# get image id based on tag or commit
IMAGE_VERSION := $(or $(IMAGE_TAG),$(IMAGE_COMMIT))
IMAGE_NAME := "ghcr.io/paullockaby/test-go"
IMAGE_ID := "${IMAGE_NAME}:${IMAGE_VERSION}"

all: build

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: build
build:
	mkdir -p ./build
	go build -o ./build/testrepo ./cmd/testrepo/main.go

.PHONY: buildx
buildx:
	@echo "building multiarch image for ${IMAGE_ID}"
	docker buildx build --platform linux/amd64,linux/arm64 -t $(IMAGE_ID) .

.PHONY: push
push:
	@echo "pushing $(IMAGE_ID) with buildx"
	docker buildx build --push --platform linux/amd64,linux/arm64 -t $(IMAGE_ID) -t $(IMAGE_NAME):latest .

.PHONY: clean
clean:
	rm -rf ./build
	find . -type f -name .DS_Store -print0 | xargs -0 rm -f

.PHONY: realclean
realclean: clean
	go clean -cache
	go clean -modcache

.PHONY: format
format:
	gofmt -s -w -l .

.PHONY: lint
lint:
	pre-commit run --all-files

.PHONY: pre-commit
pre-commit:
	pre-commit install

.PHONY: bump-check
bump-check:
	cz bump --dry-run
