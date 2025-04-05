package order

import (
	"context"
	"taxiservice/rider/internal/db/repository"
)

type OrderRepository interface {
	List(ctx context.Context, userID int) ([]repository.OrderModel, error)
	CreateAndGetID(ctx context.Context, order *repository.OrderModel) (string, error)
}

type RidePriceEstimator interface {
	Estimate(ctx context.Context, lat1 float32, lng1 float32, lat2 float32, lng2 float32) (int, error)
}
