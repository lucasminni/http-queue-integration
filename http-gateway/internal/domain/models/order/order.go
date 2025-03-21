package models

import (
	models "http-gateway/internal/domain/models/item"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID         uuid.UUID     `json:"id"`
	OrderDate  time.Time     `json:"order_date"`
	Status     string        `json:"status"`
	Items      []models.Item `json:"items"`
	TotalPrice float64       `json:"total_price"`
}
