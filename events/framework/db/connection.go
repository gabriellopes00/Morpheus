package db

import (
	"database/sql"
	"events/config/env"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/golang-migrate/migrate/v4"
	pgMigration "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewPostgresDb() (*gorm.DB, error) {
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *sql.DB) error {
	driver, err := pgMigration.WithInstance(db, &pgMigration.Config{})
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
