package schemas

type BodyUpdateOrderStatus struct {
	ID     string `json:"id"`
	Status string `json:"new_status"`
}
