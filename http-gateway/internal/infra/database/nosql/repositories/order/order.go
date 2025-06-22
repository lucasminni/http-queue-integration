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

	if len(orders) == 0 {
		return nil, nil
	}

	return &orders[0], nil
}

func UpdateOrderStatus(id string, status string) (*model.Order, error) {
	order, err := GetOrderById(id)

	if err != nil {
		return nil, err
	}

	updatedOrder := order.SetNewStatus(status)

	err = nosql.Insert(ORDER_DYNAMO_TABLE, updatedOrder)

	if err != nil {
		return nil, err
	}

	return &updatedOrder, nil
}
