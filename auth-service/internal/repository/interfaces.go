package repository

import (
	"context"

	"github.com/maksroxx/DivineEye/auth-service/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, id, email, password string) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}
