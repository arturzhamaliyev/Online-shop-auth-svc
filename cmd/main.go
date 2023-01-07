package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/config"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/db"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/pb"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/repository"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/services"
	"google.golang.org/grpc"
)

// TODO: database migration
// TODO: swagger connection

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config:", err)
	}

	h, err := db.Init(context.Background(), c)
	if err != nil {
		log.Fatalln("Failed at initializing db:", err)
	}
	defer h.DB.Close(context.Background())

	repo := repository.NewAuthRepository(h)

	authServer := services.NewAuthServer(repo)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Server is ready to accept clients on port :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
