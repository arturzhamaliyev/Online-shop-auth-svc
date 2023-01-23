proto:
	protoc internal/pb/*.proto --go_out=.
	protoc internal/pb/*.proto --go-grpc_out=.
	go mod tidy

doc-c_u:
	docker-compose -f docker-compose.yaml up --build

doc-c_d:
	docker-compose -f docker-compose.yaml down

server:
	go run cmd/main.go

open_postgres:
	docker exec -it online-shop-auth-svc_postgreSQL_1 psql -U auth_svc