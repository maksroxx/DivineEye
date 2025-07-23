package repository

import (
	"context"
	"database/sql"

	"github.com/maksroxx/DivineEye/notification-service/internal/models"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPstgresRepo(db *sql.DB) Alerter {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) SaveAlert(ctx context.Context, alert models.Alert) error {
	_, err := r.db.ExecContext(ctx, `
        INSERT INTO notification_alerts (id, user_id, coin, direction, price)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (id) DO NOTHING
    `, alert.ID, alert.UserID, alert.Coin, alert.Direction, alert.Price)
	return err
}

func (r *PostgresRepo) DeleteAlert(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM notification_alerts WHERE id = $1", id)
	return err
}

func (r *PostgresRepo) GetAlertsForCoin(ctx context.Context, coin string) ([]*models.Alert, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, user_id, coin, price FROM notification_alerts WHERE coin = $1", coin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*models.Alert
	for rows.Next() {
		var a models.Alert
		if err := rows.Scan(&a.ID, &a.UserID, &a.Coin, &a.Price); err != nil {
			return nil, err
		}
		alerts = append(alerts, &a)
	}
	return alerts, nil
}

func (r *PostgresRepo) GetTriggeredAlerts(ctx context.Context, coin string, price float64) ([]*models.Alert, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, coin, price, direction 
		FROM notification_alerts 
		WHERE coin = $1 AND (
			(direction = 'above' AND price <= $2) OR
			(direction = 'below' AND price >= $2)
		)
	`, coin, price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*models.Alert
	for rows.Next() {
		var a models.Alert
		if err := rows.Scan(&a.ID, &a.UserID, &a.Coin, &a.Price, &a.Direction); err != nil {
			return nil, err
		}
		alerts = append(alerts, &a)
	}
	return alerts, nil
}
