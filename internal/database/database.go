package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func NewGormDB() (*gorm.DB, error) {
	var dsn string
	// Si existe DATABASE_URL, úsala directamente (más seguro)
	if url := os.Getenv("DATABASE_URL"); url != "" {
		dsn = url
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			env("DB_HOST", "localhost"),
			env("DB_USER", "postgres"),
			env("DB_PASSWORD", "Postgres123"),
			env("DB_NAME", "aprendiz"),
			env("DB_PORT", "5432"),
		)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}