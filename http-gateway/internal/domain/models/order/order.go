package models

type Order struct {
	Id     string  `json:"id"`
	Date   string  `json:"order_date"`
	Status string  `json:"status"`
	Price  float64 `json:"total_price"`
}
