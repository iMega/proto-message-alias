default:
	@docker run --rm -v $(CURDIR):/data -v $$GOPATH:/go -w /data \
		imega/grpcgen:1.0.0 \
		-I/data \
		-I/go/src/github.com/googleapis/googleapis \
		-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle=github.com/googleapis/googleapis/google,plugins=grpc:. /data/alias.proto

build:
	docker run --rm -v $(GOPATH):/go \
		-w /go/src/github.com/imega/proto-message-alias \
		golang:1.10-alpine go build -v -o protoc-gen-alias
