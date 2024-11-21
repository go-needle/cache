package cache

import (
	"github.com/go-needle/cache/alg"
	"sync"
	"time"
)

type Cache interface {
	Add(key string, value []byte)
	Get(key string) ([]byte, bool)
}

type LRUCache struct {
	mu              sync.Mutex
	cache           *alg.LRU
	cacheBytes      int64
	maxSurvivalTime time.Duration
}

// NewLRU creates a new cache struct and use the LRU algorithm
func NewLRU(cacheBytes int64, maxSurvivalTime time.Duration) *LRUCache {
	return &LRUCache{maxSurvivalTime: maxSurvivalTime, cacheBytes: cacheBytes}
}

// Add is safe for concurrent access.
func (c *LRUCache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		c.cache = alg.NewLRU(c.cacheBytes)
	}
	c.cache.Add(key, value)
	go func() {
		time.Sleep(c.maxSurvivalTime)
		c.mu.Lock()
		defer c.mu.Unlock()
		c.cache.Delete(key)
	}()
}

// Get is safe for concurrent access.
func (c *LRUCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		return nil, false
	}
	if v, ok := c.cache.Get(key); ok {
		return v, ok
	}
	return nil, false
}

type FIFOCache struct {
	mu              sync.RWMutex
	cache           *alg.FIFO
	cacheBytes      int64
	maxSurvivalTime time.Duration
}

// NewFIFO creates a new cache struct and use the FIFO algorithm
func NewFIFO(cacheBytes int64, maxSurvivalTime time.Duration) *LRUCache {
	return &LRUCache{maxSurvivalTime: maxSurvivalTime, cacheBytes: cacheBytes}
}

// Add is safe for concurrent access.
func (c *FIFOCache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		c.cache = alg.NewFIFO(c.cacheBytes)
	}
	c.cache.Add(key, value)
	go func() {
		time.Sleep(c.maxSurvivalTime)
		c.mu.Lock()
		defer c.mu.Unlock()
		c.cache.Delete(key)
	}()
}

// Get is safe for concurrent access.
func (c *FIFOCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.cache == nil {
		return nil, false
	}
	if v, ok := c.cache.Get(key); ok {
		return v, ok
	}
	return nil, false
}
