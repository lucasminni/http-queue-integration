package services

import (
	model "http-gateway/internal/domain/models/order"
	schema "http-gateway/internal/domain/schemas/order"
	repository "http-gateway/internal/infra/database/nosql/repositories/order"

	"github.com/dromara/carbon/v2"
	uuid "github.com/satori/go.uuid"
)

func CreateOrder(order model.Order) (*model.Order, error) {
	order.Id = uuid.NewV4().String()
	order.Date = carbon.Now().ToString()
	order.Status = "processing"

	result, err := repository.CreateOrder(order)

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
