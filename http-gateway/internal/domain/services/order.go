package services

import (
	model "http-gateway/internal/domain/models/order"
	nosql "http-gateway/internal/infra/database/nosql/repositories/order"
)

func CreateOrder(price float64) (*model.Order, error) {
	order := model.Order{}
	order = order.SetNewOrder(price)
	result, err := nosql.CreateOrder(order)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetOrders() ([]model.Order, error) {
	order, err := nosql.GetOrders()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetOrderById(id string) (*model.Order, error) {
	order, err := nosql.GetOrderById(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func UpdateOrderStatus(id string, status string) (*model.Order, error) {
	updatedOrder, err := nosql.UpdateOrderStatus(id, status)

	if err != nil {
		return nil, err
	}

	return updatedOrder, nil
}

func DeleteOrder(id string) error {

	if err := nosql.DeleteOrder(id); err != nil {
		return err
	}

	return nil
}

func DeleteOderByIdList(ids []string) error {

	if err := nosql.DeleteOrderByIdList(ids); err != nil {
		return err
	}

	return nil
}
