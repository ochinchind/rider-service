package ride

import (
	"context"
	"github.com/redis/go-redis/v9"
	"ride-service/example/internal/db/repository"
	"time"
)

type RideService struct {
	rideRepository RideRepository
	conn           *redis.Client
}

func NewRideService(rideRepository RideRepository, conn *redis.Client) RideService {
	return RideService{
		rideRepository: rideRepository,
		conn:           conn,
	}
}

func (s *RideService) TrackOrder(ctx context.Context, id string, t time.Time, latitude float32, longitude float32) error {
	err := s.rideRepository.TrackPoint(ctx, id, repository.Location{
		CreatedAt: t,
		Latitude:  latitude,
		Longitude: longitude,
	})

	if err != nil {
		return err
	}

	return s.conn.XAdd(ctx, &redis.XAddArgs{
		Stream: "tracks",
		MaxLen: 1000,
		Approx: true,
		Values: map[string]interface{}{
			"lat":  latitude,
			"lon":  longitude,
			"id":   id,
			"time": t.Unix(),
		},
	}).Err()
}

func (s *RideService) GetTrack(ctx context.Context, id string) ([]Location, error) {
	track, err := s.rideRepository.GetTrack(ctx, id)
	if err != nil {
		return nil, err
	}

	locations := make([]Location, len(track))
	for i, t := range track {
		locations[i] = Location{
			CreatedAt: t.CreatedAt,
			Latitude:  t.Latitude,
			Longitude: t.Longitude,
		}
	}
	return locations, nil
}
