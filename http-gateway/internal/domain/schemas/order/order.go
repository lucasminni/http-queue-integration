package schemas

type BodyUpdateOrderStatus struct {
	ID     string `json:"id"`
	Status string `json:"new_status"`
}

type BodyNewOrder struct {
	Price float64 `json:"total_price"`
}

type BodyDeleteOrderList struct {
	IDList []string `json:"ids"`
}
