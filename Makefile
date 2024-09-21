test:
	go test -v -cover -race ./...
.PHONY: test

install-protoc:
	sudo apt update
	sudo apt install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
.PHONY: install-protoc

generate-api:
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		--proto_path=api/v1 \
		api/v1/*.proto
.PHONY: generate-api
