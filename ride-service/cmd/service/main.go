package main

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"ride-service/example/internal/config"
	"ride-service/example/internal/db/repository"
	ride_order "ride-service/example/internal/generated/proto/ride.order"
	"ride-service/example/internal/handlers"
	"ride-service/example/internal/logger"
	"ride-service/example/internal/otel"
	"ride-service/example/internal/services/ride"
	"syscall"
	"time"
)

func main() {
	log := logger.New()
	cfg, err := config.FromEnv()
	if err != nil {
		log.WithError(err, "get cfg")
		os.Exit(1)
	}

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: "",
		Protocol: 3,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.WithError(err, "redis ping")
		os.Exit(1)
	}
	defer rdb.Close()

	rdbBroker := redis.NewClient(&redis.Options{
		Addr:     cfg.BrokerURL,
		Password: "",
		Protocol: 3,
	})
	if err := rdbBroker.Ping(ctx).Err(); err != nil {
		log.WithError(err, "redis broker ping")
		os.Exit(1)
	}
	defer rdbBroker.Close()

	rideRepository := repository.NewRideRepository(rdb)
	rideService := ride.NewRideService(rideRepository, rdbBroker)

	handler := handlers.NewHandler(rideService)

	s := grpc.NewServer()
	reflection.Register(s)
	ride_order.RegisterRideServer(s, handler)

	serviceName := "ride-service"
	serviceVersion := os.Getenv("SERVICE_VERSION")
	otelShutdown, err := otel.SetupOTelSDK(ctx, serviceName, serviceVersion, cfg.Env == config.ProdEnv)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Listen", "addr", cfg.ListenAddrAndPort())
		listener, err := net.Listen("tcp", cfg.ListenAddrAndPort())
		if err != nil {
			log.WithError(err, "listen")
			close(done)
			return
		}
		if err := s.Serve(listener); err != nil {
			log.WithError(err, "listen")
			close(done)
		}
	}()

	<-done
	log.Info("Listen stopped")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		cancel()
	}()

	s.GracefulStop()
	log.Info("Shutdown completed")
}
