package routes

import (
	"api-test-crud/internal/models"
	"api-test-crud/internal/services"
	"github.com/go-chi/chi/v5"
)

func LoginRouter() chi.Router {
	models.GenerateUsers() // Test users

	r := chi.NewRouter()
	loginHandler := services.LoginHandler{}
	r.Group(func(r chi.Router) {
		r.Post("/", loginHandler.AuthUser)
	})
	return r
}
