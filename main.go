package main

import (
	"api-test-crud/config"
	"api-test-crud/internal/routes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Start Config
	config.InitEnv()
	config.InitJwt()

	// Port
	port := getPort()

	// Router
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(Cors())
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	// Routes
	r.Mount("/api/v1/login", routes.LoginRouter())
	r.Mount("/api/v1/clients", routes.ClientsRouter())
	r.Mount("/api/v1/seats", routes.SeatsRouter())

	initServer(r, port)
}

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	return port
}

func Cors() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	})
}

func initServer(r *chi.Mux, port int) {
	log.Printf("Servidor iniciado en puerto %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("Error al levantar el servidor: " + err.Error())
	}
}
