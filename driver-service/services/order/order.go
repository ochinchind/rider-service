package order

import (
	"context"
	"driver-service/driver-service/internal/db/repository"
	"driver-service/driver-service/internal/logger"
	"time"
)

type OrderService struct {
	repo         OrderRepository
	driverSearch DriverSearchRepository
	log          logger.Log
}

func NewOrderService(repo OrderRepository, driverSearch DriverSearchRepository, log logger.Log) OrderService {
	return OrderService{
		repo:         repo,
		driverSearch: driverSearch,
		log:          log,
	}
}

func (s OrderService) Create(ctx context.Context, orderCreate OrderCreate) (*OrderModel, error) {
	pickupLocation := repository.PickupLocation{
		Latitude:  orderCreate.PickupLocation.Latitude,
		Longitude: orderCreate.PickupLocation.Longitude,
	}

	dropoffLocation := repository.DropoffLocation{
		Latitude:  orderCreate.DropoffLocation.Latitude,
		Longitude: orderCreate.DropoffLocation.Longitude,
	}

	order := repository.OrderModel{
		ID:              orderCreate.ID,
		PickupLocation:  pickupLocation,
		DropoffLocation: dropoffLocation,
		UserID:          orderCreate.UserID,
	}

	driverID, err := s.driverSearch.FindDriver(
		ctx,
		orderCreate.PickupLocation.Latitude,
		orderCreate.PickupLocation.Longitude,
		orderCreate.DropoffLocation.Latitude,
		orderCreate.DropoffLocation.Longitude,
	)

	if err == nil {
		order.DriverID = &driverID
	}

	err = s.repo.Create(ctx, &order)
	if err != nil {
		return nil, err
	}

	if order.DriverID == nil {
		go func() {
			attempts := 0
			for attempts < 10 {
				attempts++
				time.Sleep(time.Second * 1)
				driverID, err := s.driverSearch.FindDriver(
					ctx,
					orderCreate.PickupLocation.Latitude,
					orderCreate.PickupLocation.Longitude,
					orderCreate.DropoffLocation.Latitude,
					orderCreate.DropoffLocation.Longitude,
				)
				if err != nil {
					s.log.Error("failed to find driver", err)
					continue
				}
				ok, err := s.repo.AssignDriver(context.Background(), orderCreate.ID, driverID)
				if err != nil {
					s.log.WithError(err, "failed to assign driver", "orderID", orderCreate.ID)
					continue
				}

				if !ok {
					s.log.Warning("already assigned driver", "orderID", orderCreate.ID)
					return
				}

				s.log.Info("driver assigned", "orderID", orderCreate.ID, "driverID", driverID)
			}
		}()
	}

}
