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
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "calendar" (
		id SERIAL PRIMARY KEY,
		user TEXT NOT NULL,
		time TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
