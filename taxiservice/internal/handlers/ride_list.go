package handlers

import (
	"encoding/json"
	"net/http"
	"taxiservice/rider/internal/generated/schema"
)

func (h *RideImpl) GetOrders(w http.ResponseWriter, r *http.Request, params rider.rider) {
	if params.XUserId <= 0 {
		writeAuthError(w)
		return
	}

	orders, err := h.orderService.List(r.Context(), params.XUserId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	resultOrders := make([]rider.Order, 0, len(orders))
	for _, order := range orders {
		resultOrders = append(resultOrders, rider.Order{
			CreatedAt: order.CreatedAt.String(),
			DropoffLocation: rider.Location{
				Latitude:  order.DropoffLocation.Latitude,
				Longitude: order.DropoffLocation.Longitude,
			},
			PickupLocation: rider.Location{
				Latitude:  order.PickupLocation.Latitude,
				Longitude: order.PickupLocation.Longitude,
			},
			Id:         order.ID,
			TotalPrice: order.TotalPrice,
		})
	}

	_ = json.NewEncoder(w).Encode(&orders)
}
