package services

import (
	model "http-gateway/internal/domain/models/order"
	schema "http-gateway/internal/domain/schemas/order"
	nosql "http-gateway/internal/infra/database/nosql/repositories/order"
)

func CreateOrder(order model.Order) (*model.Order, error) {
	order = model.Order.SetNewOrder(order)
	result, err := nosql.CreateOrder(order)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetOrders() ([]model.Order, error) {
	return nil, nil
}

func UpdateOrder(schema.BodyUpdateOrder) (*model.Order, error) {
	return nil, nil
}

func DeleteOrder(id string) error {
	return nil
}
