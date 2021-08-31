package db

import (
	"accounts/config/env"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
