package auth

import (
	"context"

	"github.com/maksroxx/DivineEye/gateway/internal/grpcclient"
	pb "github.com/maksroxx/DivineEye/gateway/proto-clients/auth"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (bool, string, error)
}

type authService struct {
	client *grpcclient.AuthClient
}

func NewAuthService(client *grpcclient.AuthClient) AuthService {
	return &authService{client: client}
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	resp, err := s.client.Login(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

func (s *authService) Register(ctx context.Context, email, password string) (string, error) {
	resp, err := s.client.Register(ctx, &pb.RegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

func (s *authService) ValidateToken(ctx context.Context, token string) (bool, string, error) {
	resp, err := s.client.ValidateToken(ctx, &pb.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		return false, "", err
	}
	return resp.Valid, resp.UserId, nil
}
