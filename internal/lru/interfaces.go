package lru

type LRUCacheInterface interface {
	Add(key, value string) bool
	Get(key string) (value string, ok bool)
	Remove(key string) (ok bool)
}
