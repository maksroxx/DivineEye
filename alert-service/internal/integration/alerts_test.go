package integration

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/maksroxx/DivineEye/alert-service/internal/models"
	"github.com/maksroxx/DivineEye/alert-service/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDBDSN = "postgres://myuser:mypassword@localhost:5432/mydatabase_test?sslmode=disable"

func cleanupTables(t *testing.T, db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE alerts, users RESTART IDENTITY CASCADE;")
	require.NoError(t, err)
}

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBDSN)
	require.NoError(t, err)
	LoadFixtures(t, db)

	t.Cleanup(func() {
		cleanupTables(t, db)
	})

	return db
}

func TestGetAlerts_WithFixtures(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewPostgresRepo(db)

	alerts, err := repo.GetAll(context.Background(), "11111111-1111-1111-1111-111111111111")
	require.NoError(t, err)
	require.Len(t, alerts, 2)

	assert.Equal(t, "BTC", alerts[0].Coin)
	assert.Equal(t, "ETH", alerts[1].Coin)
}

func TestGetAlert_WithFixtures(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewPostgresRepo(db)

	alert, err := repo.GetById(context.Background(), "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")
	require.NoError(t, err)
	require.Equal(t, "ETH", alert.Coin)
}

func TestCreateAlert(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewPostgresRepo(db)

	newAlert := &models.Alert{
		ID:     "dddddddd-dddd-dddd-dddd-dddddddddddd",
		UserID: "11111111-1111-1111-1111-111111111111",
		Coin:   "ADA",
		Price:  0.50,
	}

	err := repo.Create(context.Background(), newAlert)
	require.NoError(t, err)

	alerts, err := repo.GetAll(context.Background(), "11111111-1111-1111-1111-111111111111")
	require.NoError(t, err)
	require.Len(t, alerts, 3)

	coins := []string{alerts[0].Coin, alerts[1].Coin, alerts[2].Coin}
	assert.Contains(t, coins, "ADA")
}

func TestDeleteAlert(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewPostgresRepo(db)

	err := repo.Delete(context.Background(), "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	require.NoError(t, err)

	alerts, err := repo.GetAll(context.Background(), "11111111-1111-1111-1111-111111111111")
	require.NoError(t, err)
	assert.Len(t, alerts, 1)
}

func TestCascadeDeleteUser(t *testing.T) {
	db := setupTestDB(t)

	_, err := db.Exec(`DELETE FROM users WHERE id = '11111111-1111-1111-1111-111111111111'`)
	require.NoError(t, err)

	repo := repository.NewPostgresRepo(db)
	alerts, err := repo.GetAll(context.Background(), "11111111-1111-1111-1111-111111111111")
	require.NoError(t, err)
	assert.Len(t, alerts, 0)
}
