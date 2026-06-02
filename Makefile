VERSION := $(shell git describe --tags --always 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)
CODEGRAPH_VERSION := v0.9.7

.PHONY: build vet fmt test hooks cross clean e2e-codegraph

build:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/codelo ./cmd/codelo

vet:
	go vet ./...

fmt:
	gofmt -w .

test:
	go test ./...

cross:
	@mkdir -p dist
	@for p in darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64; do \
		os=$${p%/*}; arch=$${p#*/}; ext=; [ $$os = windows ] && ext=.exe; \
		echo "build $$os/$$arch"; \
		CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -ldflags "$(LDFLAGS)" -o dist/codelo-$$os-$$arch$$ext ./cmd/codelo; \
	done

clean:
	rm -rf bin dist