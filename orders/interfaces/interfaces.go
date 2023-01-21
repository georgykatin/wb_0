package interfaces

import (
	"context"
	"wb/config"
)

type CacheRepo interface {
	GetSizeCache() int
	SetCache(ctx context.Context, OrderUid *config.Order) error
	GetByUid(ctx context.Context, OrderUid config.Order) (*config.Order, error)
	DeleteCache(ctx context.Context, OrderUid string)
}

type DbRepo interface {
	BatchCreate(ctx context.Context, Order config.Order) error
	CreateOrder(ctx context.Context, Order config.Order) (*config.Order, error)
	GetByUid(ctx context.Context, OrderUid string) (*config.Order, error)
	GetAllByUid(ctx context.Context, OrderUid string) (*config.Order, error)
	GetOrderAmount(ctx context.Context) (int, error)
	GetOrderUid(ctx context.Context) ([]string, error)
}
