package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=http_golang_test password=password dbname=http_golang_test sslmode=disable"
	}

	os.Exit(m.Run())
}
