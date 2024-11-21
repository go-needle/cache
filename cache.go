package cache

import (
	"github.com/go-needle/cache/alg"
	"sync"
	"time"
)

type Cache struct {
	mu              sync.Mutex
	cache           alg.Alg
	cacheBytes      int64
	maxSurvivalTime time.Duration
}

func New(cacheBytes int64, maxSurvivalTime time.Duration, alg alg.Alg) *Cache {
	return &Cache{cache: alg, maxSurvivalTime: maxSurvivalTime, cacheBytes: cacheBytes}
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Add(key, value)
	go func() {
		time.Sleep(c.maxSurvivalTime)
		if !c.cache.Exist(key) {
			return
		}
		c.mu.Lock()
		defer c.mu.Unlock()
		c.cache.Delete(key)
	}()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.cache.Get(key); ok {
		return v, ok
	}
	return nil, false
}
