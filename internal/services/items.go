package services

import (
	"api-test-crud/internal/models"
	"api-test-crud/internal/utils"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type CrudHandler struct{}

func (h *CrudHandler) ListItems(w http.ResponseWriter, r *http.Request) {
	utils.OkResponse(w, r, "", models.GetItems())
}

func (h *CrudHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	item, ok := models.GetItem(itemID)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Registro con ID "+itemID+" no existe")
		return
	}

	utils.OkResponse(w, r, "", item)
}

func (h *CrudHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	itemIsValid := parseItem(&item, r.Body)
	if !itemIsValid {
		utils.BadRequestResponse(w, r, "Error: Datos inválidos")
		return
	}

	newItem := models.CreateItem(&item)
	utils.CreatedResponse(w, r, "Item creado", newItem)
}

func (h *CrudHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	itemID := chi.URLParam(r, "id")

	itemIsValid := parseItem(&item, r.Body)
	if !itemIsValid {
		utils.BadRequestResponse(w, r, "Error: Datos inválidos")
		return
	}

	item.ID = itemID
	newItem, ok := models.UpdateItem(&item)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Registro con ID "+item.ID+" no existe")
		return
	}

	utils.OkResponse(w, r, "Item actualizado", newItem)
}

func (h *CrudHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	item, ok := models.DeleteItem(itemID)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Registro con ID "+itemID+" no existe")
		return
	}

	utils.OkResponse(w, r, "Item eliminado", item)
}

// Helpers

func parseItem(item *models.Item, body io.ReadCloser) bool {
	if body == nil {
		return false
	}

	if err := json.NewDecoder(body).Decode(item); err != nil {
		return false
	}

	itemIsValid := isItemPayloadValid(*item)
	return itemIsValid
}

func isItemPayloadValid(item models.Item) bool {
	if item.Name == "" {
		return false
	}
	return true
}
