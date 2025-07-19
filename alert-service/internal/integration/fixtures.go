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
