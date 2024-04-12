package main

import (
	"api-test-crud/config"
	"api-test-crud/internal/routes"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func main() {
	// Start Config
	config.InitEnv()
	config.InitJwt()

	// Flags
	port := flag.Int("port", 8080, "Número de puerto para el servidor")
	flag.Parse()

	// Router
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(Cors())
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	// Routes
	r.Mount("/api/v1/items", routes.CrudRouter())
	r.Mount("/api/v1/login", routes.LoginRouter())

	// Init server
	initServer(r, *port)
}

func Cors() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Max value
	})
}

func initServer(r *chi.Mux, port int) {
	log.Printf("Servidor iniciado en puerto %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("Error al levantar el servidor: " + err.Error())
	}
}
