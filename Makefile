proto:
	protoc pkg/pb/*.proto --go_out=.
	protoc pkg/pb/*.proto --go-grpc_out=.
	go mod tidy

docker-compose_up:
	docker-compose -f docker-compose.yaml up --build

docker-compose_down:
	docker-compose -f docker-compose.yaml down

server:
	go run cmd/main.go

open_postgres:
	docker exec -it online-shop-auth-svc_postgreSQL_1 psql -U auth_svc