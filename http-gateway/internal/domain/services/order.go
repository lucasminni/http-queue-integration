package services

import (
	model "http-gateway/internal/domain/models/order"
	schema "http-gateway/internal/domain/schemas/order"
)

func CreateOrder(model.Order) (*model.Order, error) {
	return nil, nil
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
