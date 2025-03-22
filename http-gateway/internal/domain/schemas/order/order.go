package schemas

import (
	"time"
)

type BodyNewOrder struct {
	OrderDate  time.Time `json:"order_date"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
}

type BodyUpdateOrder struct {
	OrderDate  time.Time `json:"new_order_date"`
	Status     string    `json:"new_status"`
	TotalPrice float64   `json:"new_total_price"`
}
