PROTO := $(shell find api -name "*.proto")
GO := $(shell find . -name "*.go")

proto: $(PROTO)
	protoc -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate -I . $(PROTO) --go_out=plugins=grpc:. --validate_out="lang=go:."

.PHONY: clean

run-feeds: bin/feeds
	./bin/feeds

bin/feeds: proto $(GO)
	go build -o bin/feeds cmd/main.go

docker.build:
	docker build -f docker/feeds-framework/Dockerfile . -t feeds-framework:latest
	docker-compose up