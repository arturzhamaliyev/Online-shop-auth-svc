package services

import (
	"context"
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/models"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/pb"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/repository"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/utils"
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

	err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	user.Id, err = s.repo.Create(ctx, user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Id:     user.Id,
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *Server) Validate(context.Context, *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, nil
}
