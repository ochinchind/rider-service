package handlers

import (
	"context"
	driver_order "driver-service/driver-service/internal/generated/proto/driver.order"
	"driver-service/driver-service/internal/services/order"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type Handler struct {
	driver_order.UnimplementedOrderServer
	orderService OrderService
}

func NewHandler(orderService OrderService) *Handler {
	return &Handler{
		orderService: orderService,
	}
}

func (h Handler) StartOrder(ctx context.Context, req *driver_order.StartOrderRequest) (*driver_order.StartOrderResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("start order")
	if time.Now().Unix()%2 == 0 {
		time.Sleep(time.Millisecond * 500)
	}
	driverOrder, err := h.orderService.Create(ctx, order.OrderCreate{
		ID:        req.Id,
		CreatedAt: req.CreatedAt.AsTime(),
		PickupLocation: order.Location{
			Latitude:  req.PointA.Latitude,
			Longitude: req.PointA.Longitude,
		},
		DropoffLocation: order.Location{
			Latitude:  req.PointB.Latitude,
			Longitude: req.PointB.Longitude,
		},
		UserID: req.UserID,
	})
	if err != nil {
		span.SetStatus(codes.Error, "failed start order")
		span.RecordError(err)
		return nil, err
	}
	return &driver_order.StartOrderResponse{DriverId: &driverOrder.DriverID}, nil
}
