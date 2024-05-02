package services

import (
	"api-test-crud/internal/models"
	"api-test-crud/internal/utils"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strings"
)

type ClientsHandler struct{}

func (h *ClientsHandler) ListClients(w http.ResponseWriter, r *http.Request) {
	utils.OkResponse(w, r, "", models.GetClients())
}

func (h *ClientsHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")
	client, ok := models.GetClient(clientID)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Registro con ID "+clientID+" no existe")
		return
	}

	utils.OkResponse(w, r, "", client)
}

func (h *ClientsHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	clientIsValid := parseClient(&client, r.Body)
	if !clientIsValid {
		utils.BadRequestResponse(w, r, "Error: Datos inválidos")
		return
	}

	newClient := models.CreateClient(&client)
	utils.CreatedResponse(w, r, "Cliente creado", newClient)
}

func (h *ClientsHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	clientID := chi.URLParam(r, "id")

	clientIsValid := parseClient(&client, r.Body)
	if !clientIsValid {
		utils.BadRequestResponse(w, r, "Error: Datos inválidos")
		return
	}

	client.ID = clientID
	newClient, ok := models.UpdateClient(&client)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Registro con ID "+client.ID+" no existe")
		return
	}

	utils.OkResponse(w, r, "Cliente actualizado", newClient)
}

func (h *ClientsHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")
	client, ok := models.DeleteClient(clientID)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Registro con ID "+clientID+" no existe")
		return
	}
	deletedSeats := models.DeleteSeatsFromClient(clientID)
	utils.OkResponse(w, r, "Cliente eliminado", map[string]interface{}{"client": client, "seats": deletedSeats})
}

// Helpers

func parseClient(client *models.Client, body io.ReadCloser) bool {
	if body == nil {
		return false
	}

	if err := json.NewDecoder(body).Decode(client); err != nil {
		return false
	}

	clientIsValid := isClientPayloadValid(*client)
	return clientIsValid
}

func isClientPayloadValid(client models.Client) bool {
	if strings.TrimSpace(client.Name) == "" || strings.TrimSpace(client.Rut) == "" || strings.TrimSpace(client.Birthdate) == "" {
		return false
	}
	return true
}
