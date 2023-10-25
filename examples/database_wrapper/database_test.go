package database_wrapper

import (
	"testing"
)

func mustCreateDatabase(t *testing.T) *Database {
	db, err := NewDatabase(func(o *NewDatabaseOptions) {
		o.ConnectionString = "file:database_test?mode=memory&cache=shared"
		o.RunMigrations = true
	})
	if err != nil {
		t.Fatalf("unable to create database: %s", err.Error())
		return nil
	}

	return db
}

func TestRunsMigrationsCorrectly(t *testing.T) {
	db := mustCreateDatabase(t)

	// Ensure all migrations have been ran insofar as they exist in the migrations table.
	var n int
	if err := db.db.QueryRow("SELECT COUNT(*) from migrations").Scan(&n); err != nil {
		t.Fatalf("could not query migrations table: %s", err.Error())
	} else if n != 2 {
		t.Fatalf("migrations count err, wanted 2, got: %d", n)
		return
	}

	// Ensure all migrations have taken effect on the database itself.
	var nu int
	if err := db.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&nu); err != nil {
		t.Fatalf("could not query users table: %s", err.Error())
	} else if nu != 8 {
		t.Fatalf("users count err, wanted 8, got: %d", n)
		return
	}
}

func TestDoesNotRepeatMigrations(t *testing.T) {
	mustCreateDatabase(t)
	mustCreateDatabase(t)
}
