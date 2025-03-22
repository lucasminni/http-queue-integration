package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID         uuid.UUID `json:"id"`
	OrderDate  time.Time `json:"order_date"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
}
