package mango

import "sync"

type SyncMemCache[K comparable, V any] struct {
	cache map[K]V
	mu    *sync.RWMutex
}

func NewSyncCache[K comparable, V any]() *SyncMemCache[K, V] {
	return &SyncMemCache[K, V]{
		cache: make(map[K]V),
		mu:    &sync.RWMutex{},
	}
}

func (c *SyncMemCache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.cache[key]
	return v, ok
}

func (c *SyncMemCache[K, V]) Put(key K, val V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = val
}

func (c *SyncMemCache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}
