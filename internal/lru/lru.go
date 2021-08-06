package lru

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	len      int
	storage  map[string]*list.Element
	priority *list.List
	mu       sync.RWMutex
}

type cacheItem struct {
	Key   string
	Value string
}

func NewLRUCache(n int) LRUCache {
	storage := make(map[string]*list.Element, n)
	return LRUCache{len: n, storage: storage, priority: list.New()}
}

func (c *LRUCache) Add(key, value string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.storage[key]; ok {
		c.priority.MoveToFront(el)
		el.Value.(*cacheItem).Value = value

		return false
	}

	if c.len == len(c.storage) {
		el := c.priority.Back()
		if el != nil {
			c.priority.Remove(el)
			c.Remove(el.Value.(*cacheItem).Key)
		}
	}

	el := c.priority.PushFront(&cacheItem{Key: key, Value: value})
	c.storage[key] = el

	return true
}

func (c *LRUCache) Get(key string) (value string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	el, ok := c.storage[key]

	if ok {
		c.priority.MoveToFront(el)
		value = el.Value.(*cacheItem).Value
	}

	return
}

func (c *LRUCache) Remove(key string) (ok bool) {
	_, ok = c.storage[key]

	if ok {
		delete(c.storage, key)
	}

	return
}

func (c *LRUCache) GetInternalStore() map[string]*list.Element {
	return c.storage
}
