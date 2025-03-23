package sql

import (
	model "http-gateway/internal/domain/models/order"
	dynamo "http-gateway/internal/infra/database/nosql"
)

func CreateOrder(order model.Order) (*model.Order, error) {
	err := dynamo.Insert("orders", order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetOrders() ([]model.Order, error) {
	return nil, nil
}
