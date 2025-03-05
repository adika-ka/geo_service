package main

import (
	"fmt"
	"geo_service/internal/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Post("/api/address/search", handlers.SearchHandler)
	r.Post("/api/address/geocode", handlers.GeocodeHandler)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
