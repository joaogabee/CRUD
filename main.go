package main

import (
	"CRUD/src/configs"
	"CRUD/src/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(configs.GetServerPort(), r)
}
