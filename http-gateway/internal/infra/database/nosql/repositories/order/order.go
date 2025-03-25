package nosql

import (
	model "http-gateway/internal/domain/models/order"
	"http-gateway/internal/infra/database/nosql"
)

const (
	ORDER_DYNAMO_TABLE = "orders"
)

func CreateOrder(order model.Order) (*model.Order, error) {
	err := nosql.Insert(ORDER_DYNAMO_TABLE, order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetOrders() ([]model.Order, error) {
	orders, err := nosql.Scan(ORDER_DYNAMO_TABLE)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderById(id string) (*model.Order, error) {
	orders, err := nosql.ScanById(ORDER_DYNAMO_TABLE, id)

	if err != nil {
		return nil, err
	}

	return &orders[0], nil
}
