package integration

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/maksroxx/DivineEye/auth-service/internal/auth"
	"github.com/maksroxx/DivineEye/auth-service/internal/config"
	"github.com/maksroxx/DivineEye/auth-service/internal/jwt"
	"github.com/maksroxx/DivineEye/auth-service/internal/logger"
	"github.com/maksroxx/DivineEye/auth-service/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T, dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err)

	cleanupTables(t, db)
	LoadFixtures(t, db)

	t.Cleanup(func() {
		cleanupTables(t, db)
	})

	return db
}

func TestAuthService_Integration(t *testing.T) {
	cfg, err := config.Load("../../config/config.yaml")
	require.NoError(t, err)
	log := logger.NewZapLogger()
	defer log.Sync()

	testDBDSN := "postgres://myuser:mypassword@localhost:5432/mydatabase_test?sslmode=disable"

	jwt.Init(cfg.JWT.Secret)
	db := setupTestDB(t, testDBDSN)
	repo := repository.NewPostgresRepo(db)
	svc := auth.NewAuthService(repo, log)

	// register test
	token, err := svc.Register(context.Background(), "newuser@example.com", "supersecret")
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// get by email test
	user, err := repo.GetByEmail(context.Background(), "newuser@example.com")
	require.NoError(t, err)
	assert.Equal(t, "newuser@example.com", user.Email)

	// login test
	token, err = svc.Login(context.Background(), "newuser@example.com", "supersecret")
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// validate test
	ok, userId := jwt.Validate(token)
	assert.True(t, ok)
	assert.NotEmpty(t, userId)

	fixtureUser, err := repo.GetByEmail(context.Background(), "fixtureuser@example.com")
	require.NoError(t, err)
	assert.Equal(t, "fixtureuser@example.com", fixtureUser.Email)
}
