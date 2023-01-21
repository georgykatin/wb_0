package orders

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"wb/cache"
	"wb/config"
)

type OrderCache struct {
	Data *cache.Cache
}

func (Oc *OrderCache) GetSizeCache() int {
	return len(Oc.Data.CacheStorage)
}

func OrderCacheRepo(Data *cache.Cache) *OrderCache {
	data := make(map[string]string)
	Data.CacheStorage = data

	return &OrderCache{Data: Data}
}

func (Oc *OrderCache) SetCache(ctx context.Context, Order *config.Order) error {
	OBytes, err := json.Marshal(Order)
	if err != nil {
		return errors.WithMessage(err, "Error json marshal cacherepo")
	}
	Oc.Data.InsertCache(Order.OrderUid, string(OBytes))
	log.Printf("cache length:")
	log.Printf(string(len(Oc.Data.CacheStorage)))
	return nil
}

func (Oc *OrderCache) GetByUid(ctx context.Context, OrderUid config.Order) (*config.Order, error) {
	res, err := Oc.Data.GetCache("OrderUid")
	if err != nil {
		return nil, errors.WithMessage(err, "Error get cache cacherepo  ")
	}
	var result config.Order
	if err := json.Unmarshal([]byte(res), &result); err != nil {
		return nil, errors.WithMessage(err, "Error json unmarshal cacherepo")

	}

	return &result, nil
}
func (Oc *OrderCache) DeleteCache(ctx context.Context, OrderUid string) {
	Oc.Data.DeleteCache(OrderUid)
}
