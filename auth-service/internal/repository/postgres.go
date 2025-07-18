package repository

import (
	"context"
	"database/sql"

	"github.com/maksroxx/DivineEye/auth-service/internal/models"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) UserRepository {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) CreateUser(ctx context.Context, id, email, password string) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`,
		id, email, password,
	)
	return err
}

func (r *PostgresRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	var u models.User
	if err := row.Scan(&u.Id, &u.Email, &u.Password); err != nil {
		return nil, err
	}
	return &u, nil
}
