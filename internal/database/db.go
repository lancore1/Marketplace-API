package database

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Compiler directive
//
//go:embed migration/*sql
var MigrationFS embed.FS

type Migration struct {
	srcDriver source.Driver
}

func (m *Migration) newMigration(sqlFiles embed.FS, migrationPath string) (*Migration, error) {
	// Create a new iofs driver instance.
	driver, err := iofs.New(sqlFiles, migrationPath)
	if err != nil {
		return nil, fmt.Errorf("unable to create iofs driver instance: %w", err)
	}

	return &Migration{
		srcDriver: driver,
	}, nil
}

// ApplyMigrationToDB applies the migration to the database connection.
func (m *Migration) applyMigrationToDB(db *sql.DB) error {
	// Create a new postgres driver instance.
	psqlDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create db instance: %w", err)
	}

	// Create a new migration instance with driver source and postgres driver.
	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", m.srcDriver, "psql_db", psqlDriver)
	if err != nil {
		return fmt.Errorf("unable to create migration: %w", err)
	}

	// Close the migrator when done.
	defer func() {
		migrator.Close()
	}()

	// Apply the migration
	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migration: %w", err)
	}

	return nil
}

func (m *Migration) RunMigration(migrationPath string, driverName string, connectionStr string) error {

	migrator, err := m.newMigration(MigrationFS, migrationPath)
	if err != nil {
		return fmt.Errorf("unable to create new migration at launch: %w", err)
	}

	// Get the DB instance.
	conn, err := sql.Open(driverName, connectionStr)
	if err != nil {
		return fmt.Errorf("unable with get DB instance at launch: %w", err)
	}

	defer conn.Close()

	// Apply migrations
	err = migrator.applyMigrationToDB(conn)
	if err != nil {
		return fmt.Errorf("unable with apply migration at launch: %w", err)
	}

	return nil
}
