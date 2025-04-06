package location_updater

import (
	"driver-service/driver-service/internal/logger"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type LocationUpdater struct {
	conn       *redis.Client
	repository OrderRepository
	log        logger.Logger
}

func NewLocationUpdater(conn *redis.Client, repository OrderRepository, log logger.Logger) LocationUpdater {
	return LocationUpdater{
		conn:       conn,
		repository: repository,
		log:        log,
	}
}

func (l *LocationUpdater) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			cmd := l.conn.XRead(ctx, &redis.XReadArgs{
				Streams: []string{"tracks", "$"},
				Block:   time.Second,
			})
			if cmd.Err() != nil {
				if errors.Is(cmd.Err(), redis.Nil) {
					continue
				}
				l.log.WithError(cmd.Err(), "xread")
				continue
			}
			for _, v := range cmd.Val() {
				for _, m := range v.Messages {
					jsonData, err := json.Marshal(m.Values)
					if err != nil {
						l.log.WithError(err, "json marshal")
						continue
					}
					e := LocationEvent{}
					if err = json.Unmarshal(jsonData, &e); err != nil {
						l.log.WithError(err, "json unmarshal")
						continue
					}
					err = l.repository.UpdateCurrentLocation(ctx, e.ID, Location{
						Latitude:  e.Latitude,
						Longitude: e.Longitude,
					})
					if err != nil {
						l.log.WithError(err, "update location")
						continue
					}
				}
			}
		}
	}
}
