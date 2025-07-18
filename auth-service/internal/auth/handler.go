package auth

import (
	"context"

	pb "github.com/maksroxx/DivineEye/auth-service/proto"
)

type GRPCServer struct {
	pb.UnimplementedAuthServiceServer
	Service AuthService
}

func NewGRPCServer(service AuthService) *GRPCServer {
	return &GRPCServer{Service: service}
}

func (s *GRPCServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	token, err := s.Service.Register(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{Token: token}, nil
}

func (s *GRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.Service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

func (s *GRPCServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	valid, uid := s.Service.Validate(req.Token)
	return &pb.ValidateTokenResponse{Valid: valid, UserId: uid}, nil
}
