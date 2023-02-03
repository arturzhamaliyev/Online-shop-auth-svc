package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/models"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/pb"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/repository"
	"github.com/arturzhamaliyev/Online-shop-auth-svc/internal/utils"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	repo repository.Auth
	Jwt  utils.JwtWrapper
}

func NewAuthServiceServer(repo repository.Auth, jwt utils.JwtWrapper) *Server {
	return &Server{
		repo: repo,
		Jwt:  jwt,
	}
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	_, err := s.repo.GetByEmail(ctx, req.Email)
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

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		status := http.StatusInternalServerError
		errText := "couldn't make getbyemail"
		if errors.Is(err, pgx.ErrNoRows) {
			status = http.StatusNotFound
			errText = "user not found"
		}
		return &pb.LoginResponse{
			Status: int64(status),
			Error:  errText,
		}, nil
	}

	if match := utils.CheckPasswordHash(req.Password, user.Password); !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(*user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user, err := s.repo.GetByEmail(ctx, claims.Email)
	if err != nil {
		status := http.StatusInternalServerError
		errText := "couldn't make getbyemail"
		if errors.Is(err, pgx.ErrNoRows) {
			status = http.StatusNotFound
			errText = "user not found"
		}
		return &pb.ValidateResponse{
			Status: int64(status),
			Error:  errText,
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
