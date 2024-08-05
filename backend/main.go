package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Port = "8080"

func main() {

	// Router
	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir("./web")))

	log.Printf("Сервер на порту: %s\n", Port)
	// Запуск сервера
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}
