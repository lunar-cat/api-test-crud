package services

import (
	"api-test-crud/internal/models"
	"api-test-crud/internal/utils"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type SeatsHandler struct{}

func (h *SeatsHandler) ListSeats(w http.ResponseWriter, r *http.Request) {
	utils.OkResponse(w, r, "", models.GetSeats())
}

func (h *SeatsHandler) GetSeat(w http.ResponseWriter, r *http.Request) {
	seatID := chi.URLParam(r, "id")
	seat, ok := models.GetSeatById(seatID)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Asiento con ID "+seatID+" no encontrado")
		return
	}

	utils.OkResponse(w, r, "", seat)
}

func (h *SeatsHandler) GetSeatsFromClient(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")

	_, ok := models.GetClient(clientID)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Cliente con ID "+clientID+" no existe")
		return
	}

	seats := models.GetSeatsFromClient(clientID)
	utils.OkResponse(w, r, "", seats)
}

func (h *SeatsHandler) CreateSeat(w http.ResponseWriter, r *http.Request) {
	var seat models.Seat

	seatIsValid := parseSeat(&seat, r.Body)
	if !seatIsValid {
		utils.BadRequestResponse(w, r, "Error: Datos inválidos")
		return
	}

	// Verifica que no hayan más de 32 asientos
	if len(models.GetSeats()) >= models.LimitSeats {
		utils.BadRequestResponse(w, r, "Error: Todos los asientos ocupados")
		return
	}

	// Verifica que el asiento en particular no esté ocupado
	_, occupied := models.GetSeatByPosition(seat.SeatNumber)
	if occupied {
		utils.BadRequestResponse(w, r, "Error: Asiento ya ocupado")
		return
	}

	// Verifica si el cliente existe
	_, okClient := models.GetClient(seat.ClientID)
	if !okClient {
		utils.NotFoundResponse(w, r, "Error: Cliente con ID "+seat.ClientID+" no existe")
		return
	}

	newSeat := models.CreateSeat(&seat)
	utils.CreatedResponse(w, r, "Asiento creado", newSeat)
}

func (h *SeatsHandler) DeleteSeat(w http.ResponseWriter, r *http.Request) {
	seatID := chi.URLParam(r, "id")
	seat, ok := models.DeleteSeat(seatID)
	if !ok {
		utils.NotFoundResponse(w, r, "Error: Asiento con ID "+seatID+" no encontrado")
		return
	}

	utils.OkResponse(w, r, "Asiento eliminado", seat)
}

// Helpers

func parseSeat(seat *models.Seat, body io.ReadCloser) bool {
	if body == nil {
		return false
	}

	if err := json.NewDecoder(body).Decode(&seat); err != nil {
		return false
	}

	seatIsValid := isSeatPayloadValid(*seat)
	return seatIsValid
}

func isSeatPayloadValid(seat models.Seat) bool {
	if seat.ClientID == "" {
		return false
	}

	if seat.SeatNumber <= 0 || seat.SeatNumber > models.LimitSeats {
		return false
	}

	return true
}
