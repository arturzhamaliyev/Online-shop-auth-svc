package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/config"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/db"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/pb"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/repository"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/services"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/utils"
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
	defer h.DB.Close(context.Background())

	repos := repository.NewAuthRepository(h)

	jwt := utils.JwtWrapper{
		SecretKey:      c.JWTSecretKey,
		Issuer:         "go-grpc-auth-svc",
		ExpirationTime: 24 * 30,
	}

	authServer := services.NewAuthServiceServer(repos, jwt)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	lis, err := net.Listen("tcp", ":"+c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Printf("Server is ready to accept clients on port :%s\n", c.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
