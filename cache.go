package cache

import (
	"github.com/go-needle/cache/alg"
	"sync"
	"time"
)

type Cache interface {
	Add(key string, value []byte)
	Get(key string) (ByteView, bool)
}

type LRUCache struct {
	mu              sync.Mutex
	cache           *alg.LRU
	cacheBytes      int64
	keySurvivalTime time.Duration
	once            sync.Once
}

// NewLRU creates a new cache struct and use the LRU algorithm
func NewLRU(cacheBytes int64, keySurvivalTime time.Duration) *LRUCache {
	return &LRUCache{keySurvivalTime: keySurvivalTime, cacheBytes: cacheBytes}
}

// Add is safe for concurrent access.
func (c *LRUCache) Add(key string, value []byte) {
	c.once.Do(func() {
		c.cache = alg.NewLRU(c.cacheBytes)
	})
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Add(key, value)
	go func() {
		time.Sleep(c.keySurvivalTime)
		c.mu.Lock()
		defer c.mu.Unlock()
		c.cache.Delete(key)
	}()
}

// Get is safe for concurrent access.
func (c *LRUCache) Get(key string) (ByteView, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		return ByteView{}, false
	}
	if v, ok := c.cache.Get(key); ok {
		return ByteView{v}, ok
	}
	return ByteView{}, false
}

type FIFOCache struct {
	mu              sync.RWMutex
	cache           *alg.FIFO
	cacheBytes      int64
	keySurvivalTime time.Duration
	once            sync.Once
}

// NewFIFO creates a new cache struct and use the FIFO algorithm
func NewFIFO(cacheBytes int64, keySurvivalTime time.Duration) *LRUCache {
	return &LRUCache{keySurvivalTime: keySurvivalTime, cacheBytes: cacheBytes}
}

// Add is safe for concurrent access.
func (c *FIFOCache) Add(key string, value []byte) {
	c.once.Do(func() {
		c.cache = alg.NewFIFO(c.cacheBytes)
	})
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Add(key, value)
	go func() {
		time.Sleep(c.keySurvivalTime)
		c.mu.Lock()
		defer c.mu.Unlock()
		c.cache.Delete(key)
	}()
}

// Get is safe for concurrent access.
func (c *FIFOCache) Get(key string) (ByteView, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.cache == nil {
		return ByteView{}, false
	}
	if v, ok := c.cache.Get(key); ok {
		return ByteView{v}, ok
	}
	return ByteView{}, false
}
