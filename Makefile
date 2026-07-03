HUGO ?= hugo
BLOGPOST ?= go run ./cmd/blogpost
SIGNPOSTS ?= go run ./cmd/signposts
VERIFYPOSTS ?= go run ./cmd/verifyposts

.PHONY: dev build preview clean install-cli cli post sign-posts verify-posts

dev:
	@echo "Open: http://localhost:1313/  (use http, not https)"
	$(HUGO) server -D --disableFastRender -b http://localhost:1313/ --appendPort=false

build:
	@if [ -n "$$PGP_PRIVATE_KEY" ]; then \
		echo "$$PGP_PRIVATE_KEY" | gpg --batch --import 2>/dev/null || true; \
		$(SIGNPOSTS) --repo .; \
	fi
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

sign-posts:
	$(SIGNPOSTS) --repo .

verify-posts:
	$(VERIFYPOSTS) --repo .
