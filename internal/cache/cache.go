package cache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
    mu       sync.Mutex
    entries  map[string]cacheEntry
    interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        entries:  make(map[string]cacheEntry),
        interval: interval,
    }
    go c.reapLoop()
    return c
}

func (c *Cache)reapLoop() {
ticker := time.NewTicker(c.interval)
for range ticker.C {
	c.mu.Lock()
	for key, entry := range c.entries {
		age := time.Since(entry.createdAt)
		if age > c.interval {
			delete(c.entries, key)
		}
	}
	c.mu.Unlock()
}
}