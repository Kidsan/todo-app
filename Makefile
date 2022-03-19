build:
	@CGO_ENABLED=0 go build ./cmd/...

test:
	@go test ./... -v

fmt:
	@go fmt ./...

vet:
	@go vet ./...

docker:
	@docker build . -t kidsan/todo-api:latest
	@docker tag kidsan/todo-api:latest kidsan/todo-api:$(shell git rev-parse --short --verify main)

generate:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/todo-app.proto