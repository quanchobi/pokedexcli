package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	entries map[string]cacheEntry
	mux     sync.RWMutex
	stop    chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mux:     sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = newEntry
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	entry, ok := c.entries[key]
	if !ok {
		c.mux.RUnlock()
		return nil, false
	}
	c.mux.RUnlock()
	return entry.val, true
}

func (c *Cache) Stop() {
	c.stop <- struct{}{}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-c.stop:
			return
		case t := <-ticker.C:
			c.mux.Lock()
			for key, entry := range c.entries {
				// if createdAt + interval is less than current timestamp
				if entry.createdAt.Add(interval).Compare(t) < 0 {
					delete(c.entries, key)
				}
			}
			c.mux.Unlock()
		}
	}
}
