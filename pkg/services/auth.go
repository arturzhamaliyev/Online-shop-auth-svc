package services

import (
	"context"
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/models"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/pb"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/repository"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/pkg/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	repo repository.Auth
}

func NewAuthServiceServer(repo repository.Auth) *Server {
	return &Server{repo: repo}
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if err := s.repo.GetByEmail(ctx, req.Email); err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	if err := s.repo.Create(ctx, user); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *Server) Validate(context.Context, *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, nil
}
