package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InstallDb() error {

	db, err := sql.Open("postgres", "user=tunk password=tunk host=localhost port=7755 dbname=records sslmode=disable")
	if err != nil {
		log.Fatalf("Error: Ошибка подключения к БД: %v", err)
		return err
	}
	defer db.Close()
	return nil
}
