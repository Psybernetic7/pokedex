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
	cEntry   map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cEntry:   make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cEntry[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()

	entry, ok := c.cEntry[key]

	c.mu.Unlock()

	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.cEntry {
			if time.Since(value.createdAt) > c.interval {
				delete(c.cEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
