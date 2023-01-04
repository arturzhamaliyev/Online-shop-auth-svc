proto:
	protoc pkg/pb/*.proto --go_out=.
	protoc pkg/pb/*.proto --go-grpc_out=.
	go mod tidy

postgres docker:
	docker run --name postgres-db -e POSTGRES_PASSWORD=aklpidor -d postgres 

server:
	go run cmd/main.go