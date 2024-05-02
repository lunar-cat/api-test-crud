package models

import (
	"strconv"
)

type Client struct {
	ID        string `json:"id"`
	Rut       string `json:"rut"` // Formato 11.111.111-1
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"` // Formato YYYY-MM-DD
}

// In Memory DB
var clientId = 1
var clients = []*Client{
	{
		ID:        "1",
		Rut:       "11.111.111-1",
		Name:      "Ryan Gosling",
		Birthdate: "1980-11-12",
	},
}

// Model Actions

func GetClients() []*Client {
	return clients
}

func GetClient(id string) (*Client, bool) {
	for _, client := range clients {
		if client.ID == id {
			return client, true
		}
	}
	return nil, false
}

func CreateClient(client *Client) *Client {
	clientId += 1
	newItem := &Client{
		ID:        strconv.Itoa(clientId),
		Rut:       client.Rut,
		Name:      client.Name,
		Birthdate: client.Birthdate,
	}
	clients = append(clients, newItem)
	return newItem
}

func DeleteClient(id string) (*Client, bool) {
	for i, client := range clients {
		if client.ID == id {
			clients = append(clients[:i], clients[i+1:]...)
			return client, true
		}
	}
	return nil, false
}

func UpdateClient(client *Client) (*Client, bool) {
	for i, c := range clients {
		if c.ID == client.ID {
			clients[i] = client
			return client, true
		}
	}
	return nil, false
}
