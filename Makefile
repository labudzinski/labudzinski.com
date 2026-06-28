HUGO ?= hugo
BLOGPOST ?= go run ./cmd/blogpost

.PHONY: dev build preview clean install-cli cli post

dev:
	$(HUGO) server -D --disableFastRender

build:
	$(HUGO) --minify --gc

preview: build
	$(HUGO) server --disableLiveReload --port 1313

clean:
	rm -rf public resources/_gen

install-cli:
	go build -o blogpost ./cmd/blogpost

cli: install-cli

post:
	$(BLOGPOST)
