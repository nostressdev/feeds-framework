PROTO := $(shell find api -name "*.proto")
GO := $(shell find . -name "*.go")

proto: $(PROTO)
	protoc -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate -I . $(PROTO) --go_out=plugins=grpc:. --validate_out="lang=go:."

.PHONY: clean

run-feeds: bin/feeds
	./bin/feeds

bin/feeds: proto $(GO)
	go build -o bin/feeds-framework cmd/main.go

docker.build:
	docker build -f docker/feeds-framework/Dockerfile . -t feeds-framework:latest

yc.push:
	docker tag feeds-framework:latest cr.yandex/crp3j6j9afnid0jm6v82/feeds-framework:latest
	docker push cr.yandex/crp3j6j9afnid0jm6v82/feeds-framework:latest