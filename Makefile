.PHONY: build
build:
	CGO_ENABLED=0 go build -o dist/reflect
