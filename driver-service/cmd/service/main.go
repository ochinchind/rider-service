package main

import (
	"context"
	"driver-service/driver-service/internal/config"
	"driver-service/driver-service/internal/db/repository"
	driver_order "driver-service/driver-service/internal/generated/proto/driver.order"
	"driver-service/driver-service/internal/handlers"
	"driver-service/driver-service/internal/logger"
	"driver-service/driver-service/services/driver_search"
	"driver-service/driver-service/services/order"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
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
	conn, err := pgxpool.Connect(ctx, cfg.DatabaseUrl)
	if err != nil {
		log.WithError(err, "connect to db")
		os.Exit(1)
	}
	defer conn.Close(ctx)

	driverSearch := driver_search.NewDriverSearchService()
	orderRepository := repository.NewOrderRepository(conn)
	orderService := order.NewOrderService(orderRepository, driverSearch, log)

	handler := handlers.NewHandler(orderService)

	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor(log, cfg)))
	reflection.Register(s)
	driver_order.RegisterOrderServer(s, handler)

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	s.GracefulStop()
	log.Info("Shutdown completed")
}

func loggingInterceptor(log logger.Log, cfg *config.Config) grpc.UnaryServerInterceptor {
	f := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if cfg.Env == config.LocalEnv {
			log.Info("Handle", "method", info.FullMethod, "req", req)
		}
		h, err := handler(ctx, req)
		return h, err
	}
	return f
}
