package models

import (
	"strconv"
)

type Seat struct {
	ID         string `json:"id"`
	SeatNumber int    `json:"seat_number"`
	ClientID   string `json:"client_id"`
}

const LimitSeats = 10

// In Memory DB
var seatsId = 1
var seats = []*Seat{
	{
		ID:         "1",
		SeatNumber: 1,
		ClientID:   "1",
	},
}

func GetSeats() []*Seat {
	return seats
}

func GetSeatById(id string) (*Seat, bool) {
	for _, seat := range seats {
		if seat.ID == id {
			return seat, true
		}
	}
	return nil, false
}

func GetSeatByPosition(seatNumber int) (*Seat, bool) {
	for _, seat := range seats {
		if seat.SeatNumber == seatNumber {
			return seat, true
		}
	}
	return nil, false
}

func GetSeatsFromClient(clientID string) []*Seat {
	var userSeats []*Seat
	for _, seat := range seats {
		if seat.ClientID == clientID {
			userSeats = append(userSeats, seat)
		}
	}
	return userSeats
}

func CreateSeat(seat *Seat) *Seat {
	seatsId += 1
	newSeat := &Seat{
		ID:         strconv.Itoa(seatsId),
		SeatNumber: seat.SeatNumber,
		ClientID:   seat.ClientID,
	}
	seats = append(seats, newSeat)
	return newSeat
}

func DeleteSeat(id string) (*Seat, bool) {
	for i, seat := range seats {
		if seat.ID == id {
			seats = append(seats[:i], seats[i+1:]...)
			return seat, true
		}
	}
	return nil, false
}
