package handlers

import (
	"context"
	"taxiservice/rider/internal/services/order"
)

type OrderService interface {
	Create(ctx context.Context, orderCreate order.OrderCreate) (*order.OrderModel, error)
	List(ctx context.Context, userID int) ([]order.OrderModel, error)
}
