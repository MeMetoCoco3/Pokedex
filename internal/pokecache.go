package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c.cache[key]; ok {
		return entry.val, true
	}
	return []byte{}, false
}

func (c *Cache) cleanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.cache {
		if time.Now().Sub(entry.createdAt) > 5000 {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	// .C es el unico valor que incluye el estructo Ticker
	for range ticker.C {
		c.cleanUp()
	}
}
