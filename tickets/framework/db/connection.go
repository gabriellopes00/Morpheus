package db

import (
	"database/sql"
	"fmt"
	"tickets/config/env"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewPostgresDb() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		env.DB_HOST,
		env.DB_USER,
		env.DB_PASS,
		env.DB_NAME,
		env.DB_PORT,
		env.DB_SSL_MODE,
		env.DB_TIME_ZONE,
	)

	db, err := sql.Open(env.DB_DRIVER, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance("file://framework/db/migrations", env.DB_NAME, driver)
	if err != nil {
		return err
	}

	if err = migration.Up(); err != nil {
		return err
	}

	return nil
}
