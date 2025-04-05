package order

import (
	"time"
)

type OrderModel struct {
	Id              string     `pg:"id"`
	CreatedAt       time.Time  `pg:"created_at"`
	CompletedAt     *time.Time `pg:"completed_at"`
	PickupLocation  Location   `pg:"pickup_location"`
	DropoffLocation Location   `pg:"dropoff_location"`
	TotalPrice      int        `pg:"total_price"`
	UserID          int        `pg:"user_id"`
	IdempotencyKey  string     `pg:"idempotency_key"`
}

type OrderCreate struct {
	ID              string     `pg:"id"`
	CreatedAt       time.Time  `pg:"created_at"`
	CompletedAt     *time.Time `pg:"completed_at"`
	PickupLocation  Location   `pg:"pickup_location"`
	DropoffLocation Location   `pg:"dropoff_location"`
	TotalPrice      int        `pg:"total_price"`
	UserID          int        `pg:"user_id"`
	IdempotencyKey  string     `pg:"idempotency_key"`
}

type Location struct {
	Latitude  float32 `pg:"latitude"`
	Longitude float32 `pg:"longitude"`
}
