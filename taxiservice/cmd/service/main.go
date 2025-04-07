package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yarlson/chiprom"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"taxiservice/rider/internal/config"
	"taxiservice/rider/internal/db/repository"
	"taxiservice/rider/internal/generated/proto/driver.order"
	rider "taxiservice/rider/internal/generated/schema"
	"taxiservice/rider/internal/handlers"
	"taxiservice/rider/internal/logger"
	"taxiservice/rider/internal/now_time"
	"taxiservice/rider/internal/services/driver_sender"
	"taxiservice/rider/internal/services/order"
	"taxiservice/rider/internal/services/price_estimator"
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
	conn, err := pgxpool.New(ctx, cfg.DatabaseUrl)
	if err != nil {
		log.WithError(err, "connect to db")
		os.Exit(1)
	}
	defer conn.Close()

	grpcConn, err := grpc.DialContext(
		context.Background(),
		cfg.DriverServiceLocation,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.WithError(err, "grpc connect")
		os.Exit(1)
	}

	grpcClient := driver_order.NewOrderClient(grpcConn)
	driverSenderService := driver_sender.NewDriverSenderService(grpcClient)

	priceEstimator := price_estimator.NewPriceEstimatorService()
	orderRepository := repository.NewOrderRepository(conn)
	orderService := order.NewOrderService(orderRepository, priceEstimator, now_time.Get, driverSenderService)

	handle := handlers.New(log, now_time.Get, orderService)

	r := chi.NewRouter()
	swagger, err := rider.GetSwagger()
	if err != nil {
		log.WithError(err, "get swagger")
		os.Exit(1)
	}

	r.Use(middleware.OapiRequestValidator(swagger))
	r.Use(chimiddleware.Recoverer)
	r.Use(chiprom.NewMiddleware("rider-service"))
	if cfg.Env == config.LocalEnv {
		r.Use(chimiddleware.Recoverer)
	}

	baseRouter := chi.NewRouter()
	baseRouter.Handle("/metrics", promhttp.Handler())

	rider.HandlerFromMux(handle, r)
	baseRouter.Mount("/", r)

	s := &http.Server{
		Handler: baseRouter,
		Addr:    cfg.ListenAddrAndPort(),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Listen", "addr", cfg.ListenAddrAndPort())
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err, "Listen")
			close(done)
		}
	}()

	<-done
	log.Info("Listen stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		log.Error("Shutdown error", "error", err.Error())
		os.Exit(1)
	}

	log.Info("Shutdown completed")
}
