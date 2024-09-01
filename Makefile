.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: dev
dev:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X debtrecyclingcalc.com/internal/buildinfo.GitTag=dev" \
		-o ./tmp/main ./cmd/ && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "debtrecyclingcalc.com/internal/buildinfo.GitTag=$(git describe --tags)" \
		-o ./bin/main ./cmd/
