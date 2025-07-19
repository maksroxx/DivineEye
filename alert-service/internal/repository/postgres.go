package repository

import (
	"context"
	"database/sql"

	"github.com/maksroxx/DivineEye/alert-service/internal/models"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) AlertRepository {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Create(ctx context.Context, alert *models.Alert) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO alerts (id, user_id, coin, price) VALUES ($1, $2, $3, $4)`,
		alert.ID, alert.UserID, alert.Coin, alert.Price,
	)
	return err
}

func (r *PostgresRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM alerts WHERE id = $1`, id)
	return err
}

func (r *PostgresRepo) GetAll(ctx context.Context, userId string) ([]*models.Alert, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, coin, price FROM alerts WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*models.Alert
	for rows.Next() {
		var a models.Alert
		if err := rows.Scan(&a.ID, &a.Coin, &a.Price); err != nil {
			return nil, err
		}
		a.UserID = userId
		alerts = append(alerts, &a)
	}
	return alerts, nil
}
