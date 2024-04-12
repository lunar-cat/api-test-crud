package routes

import (
	"api-test-crud/config"
	"api-test-crud/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func CrudRouter() chi.Router {

	r := chi.NewRouter()
	crudHandler := services.CrudHandler{}
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))      // Search - Verify - Validate JWT Tokens
		r.Use(jwtauth.Authenticator(config.TokenAuth)) // Verify the token itself, duration, etc.

		r.Get("/", crudHandler.ListItems)
		r.Get("/{id}", crudHandler.GetItem)
		r.Post("/", crudHandler.CreateItem)
		r.Put("/{id}", crudHandler.UpdateItem)
		r.Delete("/{id}", crudHandler.DeleteItem)
	})
	return r
}
