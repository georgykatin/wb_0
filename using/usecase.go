package Using

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"wb/config"
	"wb/orders/interfaces"
)

type UseCase struct {
	orderDB interfaces.DbRepo
	cache   interfaces.CacheRepo
}

func NewOrderUseCase(orderDB interfaces.DbRepo, cache interfaces.CacheRepo) UseCase {
	return UseCase{
		orderDB: orderDB,
		cache:   cache,
	}
}
func (order *UseCase) Create(ctx context.Context, Order config.Order) error {
	_, err := order.orderDB.CreateOrder(ctx, Order)
	if err != nil {
		return errors.WithMessage(err, "Error use case create ")
	}
	return nil
}
func (order *UseCase) BatchCreate(ctx context.Context, Order config.Order) error {
	err := order.orderDB.BatchCreate(ctx, Order)
	if err != nil {
		return errors.WithMessage(err, "Error use case batch create")
	}
	return nil
}
func (order *UseCase) GetByUid(ctx context.Context, OrderUid string) (*config.Order, error) {
	if order.cache.GetSizeCache() == 0 {
		OrderAmount, err := order.orderDB.GetOrderAmount(ctx)
		if err != nil {
			return nil, errors.WithMessage(err, "Error use case get order amount")
		}
		fmt.Printf("Order Amount : %v \n", OrderAmount)
		orderUid, err := order.orderDB.GetOrderUid(ctx)
		if err != nil {
			return nil, errors.WithMessage(err, "Error use case get order uid ")
		}
		for _, Uid := range orderUid {
			orderGetAll, err := order.orderDB.GetAllByUid(ctx, Uid)
			if err != nil {
				return nil, errors.WithMessage(err, "Error use case get all by id")
			}
			if err := order.cache.SetCache(ctx, orderGetAll); err != nil {
				log.Printf("Cache set cache : %v", err)
			}
		}
	}
	cacheUid, err := order.orderDB.GetByUid(ctx, "OrderUid")
	if err != nil {
		log.Printf("Error use case cache order get by uid")
	}
	if cacheUid != nil {
		return cacheUid, nil
	}
	orderUID, err := order.orderDB.GetAllByUid(ctx, "OrderUid")
	if err != nil {
		return nil, errors.WithMessage(err, "Error use case get by uid ")
	}
	if err := order.cache.SetCache(ctx, orderUID); err != nil {
		log.Printf("Error use case cache by uid ")
	}
	return orderUID, nil
}
