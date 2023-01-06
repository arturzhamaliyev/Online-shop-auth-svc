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
	Repo *repository.Repository
}

func NewAuthServer(repo *repository.Repository) *Server {
	return &Server{Repo: repo}
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if err := s.Repo.GetByEmail(ctx, req.Email); err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	if err := s.Repo.Create(ctx, user); err != nil {
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
