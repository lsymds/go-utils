package database_wrapper

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

// Database represents a SQLite database.
type Database struct {
	db *sql.DB
}

// NewDatabaseOptions is a set of options used to configure the database.
type NewDatabaseOptions struct {
	ConnectionString string
	RunMigrations    bool
}

// NewDatabase creates a SQLite connection and then runs any applicable migrations or seeders.
func NewDatabase(c func(*NewDatabaseOptions)) (*Database, error) {
	o := &NewDatabaseOptions{}
	c(o)

	con, err := sql.Open("sqlite3", o.ConnectionString)
	if err != nil {
		return nil, err
	}

	db := &Database{
		db: con,
	}

	// Run the migrations if they have been enabled.
	if o.RunMigrations {
		if err := db.migrate(); err != nil {
			return nil, fmt.Errorf("migrate: %w", err)
		}
	}

	return db, nil
}

// migrate migrates the database within a transaction, rolling it back and returning the error
// should any occurr.
func (d *Database) migrate() error {
	// You have to enable WAL outside of a transaction.
	if _, err := d.db.Exec("PRAGMA journal_mode = wal;"); err != nil {
		return fmt.Errorf("unable to enable wal: %w", err)
	}

	// You have to enable foreign key checks outside of a transaction.
	if _, err := d.db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return fmt.Errorf("unable to enable foreign keys: %w", err)
	}

	// Create the migrations table if it doesn't yet exist.
	if _, err := d.db.Exec("CREATE TABLE IF NOT EXISTS migrations (name TEXT PRIMARY KEY);"); err != nil {
		return fmt.Errorf("create migration table: %w", err)
	}

	// Retrieve a list of migration files to execute.
	fileNames, err := fs.Glob(migrationFS, "migrations/*.sql")
	if err != nil {
		return fmt.Errorf("globbing migration files: %w", err)
	}
	sort.Strings(fileNames)

	// Then execute them all.
	for _, fileName := range fileNames {
		if err = d.migrateFile(fileName); err != nil {
			return err
		}
	}

	return nil
}

// migrateFile runs a migration file if it hasn't been ran already.
func (d *Database) migrateFile(fileName string) error {
	// Begin a transaction for the migration.
	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("unable to start tx: %w", err)
	}

	defer tx.Rollback()

	// Check if the migration has been ran before and, if it has, return early.
	var c int
	if err := tx.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = ?", fileName).Scan(&c); err != nil {
		return err
	} else if c != 0 {
		return nil
	}

	// Read the file and execute it against the database.
	if buf, err := fs.ReadFile(migrationFS, fileName); err != nil {
		return err
	} else if _, err := tx.Exec(string(buf)); err != nil {
		return err
	}

	// Insert the record into the table.
	if _, err := tx.Exec("INSERT INTO migrations (name) VALUES (?)", fileName); err != nil {
		return err
	}

	tx.Commit()

	return nil
}
