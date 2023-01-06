package main

import (
	"context"
	"log"
	"net"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/config"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/db"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/pb"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/repository"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config:", err)
	}

	h, err := db.Init(context.Background(), c)
	if err != nil {
		log.Fatalln("Failed at initializing db:", err)
	}

	repo := repository.NewAuthRepository(h)

	server := services.NewAuthServer(repo)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
