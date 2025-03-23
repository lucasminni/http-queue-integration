package schemas

import (
	"time"

	uuid "github.com/google/uuid"
)

type BodyUpdateOrder struct {
	ID         uuid.UUID `json:"id"`
	Date       time.Time `json:"new_order_date"`
	Status     string    `json:"new_status"`
	TotalPrice float64   `json:"new_total_price"`
}
