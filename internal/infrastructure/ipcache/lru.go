package ipcache

import (
	"time"

	lru "github.com/hashicorp/golang-lru/v2/expirable"
)

type Cache[K comparable, V any] interface {
	Add(key K, value V)
	Get(key K) (V, bool)
	Purge()
}

type LRUCache[K comparable, V any] struct {
	cache *lru.LRU[K, V]
}

func NewLRU[K comparable, V any](
	size int,
	ttl time.Duration,
) *LRUCache[K, V] {

	return &LRUCache[K, V]{
		cache: lru.NewLRU[K, V](
			size,
			nil,
			ttl,
		),
	}
}

func (c *LRUCache[K, V]) Add(key K, value V) {
	c.cache.Add(key, value)
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	return c.cache.Get(key)
}

func (c *LRUCache[K, V]) Purge() {
	c.cache.Purge()
}