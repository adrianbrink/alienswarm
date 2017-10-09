.PHONY: dist
dist:
	go get github.com/mitchellh/gox
	mkdir -p dist
	gox -os="linux osx" -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

.PHONY: build
build:
	mkdir -p dist
	go build -o dist/alienswarm

.PHONY: bench
bench:
	go test github.com/eastside-eng/alienswarm/swarm -bench=.
