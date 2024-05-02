package routes

import (
	"api-test-crud/config"
	"api-test-crud/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func ClientsRouter() chi.Router {

	r := chi.NewRouter()
	clientsHandler := services.ClientsHandler{}
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator(config.TokenAuth))

		r.Get("/", clientsHandler.ListClients)
		r.Get("/{id}", clientsHandler.GetClient)
		r.Post("/", clientsHandler.CreateClient)
		r.Put("/{id}", clientsHandler.UpdateClient)
		r.Delete("/{id}", clientsHandler.DeleteClient)
	})
	return r
}
