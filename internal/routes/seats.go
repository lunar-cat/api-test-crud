package routes

import (
	"api-test-crud/config"
	"api-test-crud/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func SeatsRouter() chi.Router {

	r := chi.NewRouter()
	seatsHandler := services.SeatsHandler{}
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator(config.TokenAuth))

		r.Get("/", seatsHandler.ListSeats)
		r.Get("/{id}", seatsHandler.GetSeat)
		r.Get("/clients/{id}", seatsHandler.GetSeatsFromClient)
		r.Post("/", seatsHandler.CreateSeat)
		r.Delete("/{id}", seatsHandler.DeleteSeat)
	})
	return r
}
