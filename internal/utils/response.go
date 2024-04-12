package utils

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func OkResponse(w http.ResponseWriter, r *http.Request, Message string, Data any) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, &Response{
		true,
		Message,
		Data,
	})
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, Message string) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, &Response{
		false,
		Message,
		nil,
	})
}

func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, Message string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, &Response{
		false,
		Message,
		nil,
	})
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, Message string) {
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, &Response{
		false,
		Message,
		nil,
	})
}

func CreatedResponse(w http.ResponseWriter, r *http.Request, Message string, Data any) {
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, &Response{
		true,
		Message,
		Data,
	})
}
