package integration

import (
	"database/sql"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
)

func LoadFixtures(t *testing.T, db *sql.DB) {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../../../testdata/fixtures"),
	)
	if err != nil {
		t.Fatalf("fixtures setup failed: %v", err)
	}
	if err := fixtures.Load(); err != nil {
		t.Fatalf("fixtures load failed: %v", err)
	}
}

func cleanupTables(t *testing.T, db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")
	if err != nil {
		t.Fatalf("failed to cleanup tables: %v", err)
	}
}
