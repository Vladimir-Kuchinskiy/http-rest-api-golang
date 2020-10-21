package database

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

var testDatabaseURL = "host=localhost port=5432 user=http_golang_test password=password dbname=http_golang_test sslmode=disable"

// TestDB ...
func TestDB(t *testing.T) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", testDatabaseURL)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}
