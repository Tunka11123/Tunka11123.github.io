package db

import (
	"database/sql"
	"log"

	cfg "Tunka11123/backend/config"

	_ "github.com/lib/pq"
)

func InstallDb() error {
	dbc := new(cfg.Config)
	db, err := sql.Open("postgres", dbc.DSN())
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
