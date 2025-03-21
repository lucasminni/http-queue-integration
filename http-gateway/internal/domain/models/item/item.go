package models

import uuid "github.com/satori/go.uuid"

type Item struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
