.PHONY: build
build: build/wasm build/page

.PHONY: build/wasm
build/wasm:
	GOOS=js GOARCH=wasm go build -o src/golangci-lint-linter.wasm ./cmd/golangci-lint-linter

.PHONY: build/page
build/page:
	npm run build

.PHONY: dev
dev:
	npm run dev
