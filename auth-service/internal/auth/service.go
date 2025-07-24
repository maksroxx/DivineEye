package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/google/uuid"
	"github.com/maksroxx/DivineEye/auth-service/internal/jwt"
	"github.com/maksroxx/DivineEye/auth-service/internal/logger"
	"github.com/maksroxx/DivineEye/auth-service/internal/repository"
	"go.uber.org/zap"
)

var ErrUnauthorized = errors.New("invalid credentials")

type AuthService interface {
	Register(ctx context.Context, email, password string) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
	Validate(token string) (bool, string)
}

type authService struct {
	repo   repository.UserRepository
	logger logger.Logger
}

func NewAuthService(repo repository.UserRepository, log logger.Logger) AuthService {
	return &authService{repo: repo, logger: log}
}

func (s *authService) Register(ctx context.Context, email, password string) (string, error) {
	if _, err := s.repo.GetByEmail(ctx, email); err == nil {
		return "", errors.New("email already in use")
	}
	hash := sha256.Sum256([]byte(password))
	id := uuid.NewString()

	err := s.repo.CreateUser(ctx, id, email, hex.EncodeToString(hash[:]))
	if err != nil {
		s.logger.Error("[Service.go] Failed to create user", zap.Error(err))
		return "", err
	}
	s.logger.Info("[Service.go] User registered", zap.String("email", email))

	return jwt.Generate(id)
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", ErrUnauthorized
	}
	hash := sha256.Sum256([]byte(password))
	if user.Password != hex.EncodeToString(hash[:]) {
		return "", ErrUnauthorized
	}
	s.logger.Info("[Service.go] User logged in", zap.String("email", email))
	return jwt.Generate(user.Id)
}

func (s *authService) Validate(token string) (bool, string) {
	return jwt.Validate(token)
}
