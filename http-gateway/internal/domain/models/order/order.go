package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Id     string    `json:"id"`
	Date   time.Time `json:"order_date"`
	Status string    `json:"status"`
	Price  float64   `json:"total_price"`
}

func (order Order) SetNewOrder(price float64) Order {
	order = Order{
		Id:     uuid.NewV4().String(),
		Date:   time.Now(),
		Status: "created",
		Price:  price,
	}
	return order
}

func (order Order) SetNewStatus(status string) Order {
	order.Status = status
	return order
}
