package main

import (
	"fmt"
	"log"
	"net/http"

	cfg "Tunka11123/backend/config"

	"github.com/go-chi/chi/v5"
)

func main() {

	config := &cfg.Config{}
	// Загрузка файла .yaml
	config.ReadYaml("./config.yaml")
	fmt.Printf("%+v\n", config)
	// Роутер
	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir("./web")))

	// Запуск сервера
	log.Printf("Сервер на порту: %s\n", config.Serv.Port)
	err := http.ListenAndServe(config.Serv.Port, r)
	if err != nil {
		log.Fatal(err)
	}
}
