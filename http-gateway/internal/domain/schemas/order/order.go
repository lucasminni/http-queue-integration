package schemas

import (
	models "http-gateway/internal/domain/models/item"
	"time"
)

type BodyNewOrder struct {
	OrderDate  time.Time     `json:"order_date"`
	Status     string        `json:"status"`
	Items      []models.Item `json:"items"`
	TotalPrice float64       `json:"total_price"`
}

type BodyUpdateOrder struct {
	OrderDate  time.Time     `json:"new_order_date"`
	Status     string        `json:"new_status"`
	Items      []models.Item `json:"new_items_id_list"`
	TotalPrice float64       `json:"new_total_price"`
}
