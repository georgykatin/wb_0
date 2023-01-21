package cache

import (
	"fmt"
	"sync"
)

type Cache struct {
	mutex        sync.RWMutex
	CacheStorage map[string]string
}

func (c *Cache) GetCache(Uid string) (string, error) {
	var (
		err   error
		Order string
	)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.CacheStorage[Uid] = Order

	if Order, Presence := c.CacheStorage[Uid]; Presence {
		return Order, nil
	}
	return Order, fmt.Errorf("Cache not found %v", err)
}

func (c *Cache) DeleteCache(Uid string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.CacheStorage, Uid)
}

func (c *Cache) InsertCache(Uid string, Order string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.CacheStorage[Uid] = Order
}
