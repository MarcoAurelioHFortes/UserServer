package db

import (
	"MagicTableAPI/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresStorage() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		config.Envs.DBUsername,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBDatabase,
		config.Envs.DBSchema)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)

	}
	return db, nil
}
