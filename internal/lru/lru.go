package lru

import "sync"

type LRUCache struct {
	storage map[string]string
	mu      sync.RWMutex
}

func NewLRUCache(n int) LRUCache {
	storage := make(map[string]string, n)
	return LRUCache{storage: storage}
}

func (c *LRUCache) Add(key, value string) bool {
	return false
}

func (c *LRUCache) Get(key string) (value string, ok bool) {
	return "", false
}

func (c *LRUCache) Remove(key string) (ok bool) {
	return false
}

func (c *LRUCache) GetInternalStore() map[string]string {
	return c.storage
}
