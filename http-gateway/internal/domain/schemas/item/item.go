package schemas

type BodyNewItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type BodyUpdateItem struct {
	Name  string  `json:"new_name"`
	Price float64 `json:"new_price"`
}
