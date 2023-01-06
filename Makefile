proto:
	protoc pkg/pb/*.proto --go_out=.
	protoc pkg/pb/*.proto --go-grpc_out=.
	go mod tidy

docker_build:
	

docker-compose_up:
	docker-compose -f docker-compose.yaml up

docker-compose_down:
	docker-compose -f docker-compose.yaml down

server:
	go run cmd/main.go