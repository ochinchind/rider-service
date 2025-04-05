package handlers

import (
	"context"
	"driver-service/driver-service/services/order"
)

type OrderService interface {
	Create(ctx context.Context, orderCreate order.OrderCreate) (*order.OrderModel, error)
	UpdateCurrentLocation(ctx context.Context, id string, l order.OrderLocation) error
}
