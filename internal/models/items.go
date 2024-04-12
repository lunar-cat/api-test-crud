package models

import (
	"strconv"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// In Memory DB

var items = []*Item{
	{
		ID:   "1",
		Name: "Gato",
	},
}

// Model Actions

func GetItems() []*Item {
	return items
}

func GetItem(id string) (*Item, bool) {
	for _, item := range items {
		if item.ID == id {
			return item, true
		}
	}
	return nil, false
}

func CreateItem(item *Item) *Item {
	id := len(items) + 1
	newItem := &Item{
		ID:   strconv.Itoa(id),
		Name: item.Name,
	}
	items = append(items, newItem)
	return newItem
}

func DeleteItem(id string) (*Item, bool) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return item, true
		}
	}
	return nil, false
}

func UpdateItem(item *Item) (*Item, bool) {
	for i, it := range items {
		if it.ID == item.ID {
			items[i] = item
			return item, true
		}
	}
	return nil, false
}
