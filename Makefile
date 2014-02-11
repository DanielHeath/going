GOPATH := $(shell pwd)
export GOPATH

.PHONY: public

public:
	mkdir -p public
	rm public/*
	touch public/manifest.json
	rm -rf sample
	dashing new sample
	sprockets

bin/going: $(wildcard src/**/*.go)
	go build -o bin/going going
