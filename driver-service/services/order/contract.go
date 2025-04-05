package order

import (
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, orderCreate OrderCreate) (*OrderModel, error)
}

type DriverSearchRepository interface {
	FindDriver(ctx context.Context, lat1 float32, lng1 float32, lat2 float32, lng2 float32) (int64, error)
}
