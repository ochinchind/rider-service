package handlers

type Handler struct {
	driver_order.UnimplementedOrderServer
	orderService OrderService
}

func NewHandler(orderService OrderService) *Handler {
	return &Handler{
		orderService: orderService,
	}
}

func (h *Handler) StartOrder(ctx context.Context, req *driver_order.StartOrderRequest) (*driver_order.StartOrderResponse, error) {
	driverOrder, err := h.orderService.Create(ctx, order.OrderCreate{
		ID:        req.Id,
		CreatedAt: req.CreatedAt.AsTime(),
		PickupLocation: order.Location{
			Latitude:  req.PickupLocation.Latitude,
			Longitude: req.PickupLocation.Longitude,
		},
		DropoffLocation: order.Location{
			Latitude:  req.DropoffLocation.Latitude,
			Longitude: req.DropoffLocation.Longitude,
		},
		UserID: req.UserID,
	})

	if err != nil {
		return nil, err
	}

	return &driver_order.StartOrderResponse{DriverId: &driverOrder.DriverID}, nil
}
