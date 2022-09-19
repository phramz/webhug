BUILD_COMMIT:=$(or $(BUILD_COMMIT),$(shell git rev-parse --short HEAD))

.PHONY: vendors
vendors:
	go mod download

.PHONY: format
format:
	go fmt ./...

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint golangci-lint run

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: run
run:
	go run cmd/webhug.go

.PHONY: build
build: vendors
	mkdir -p build/bin
	CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o build/bin/webhug cmd/webhug.go

.PHONY: build-docker
build-docker:
	docker build -f Dockerfile \
		--build-arg "BUILD_COMMIT=$(BUILD_COMMIT)" \
		-t webhug \
		-t webhug:latest \
		-t webhug:$(BUILD_COMMIT) \
		 .
