package pokecache

import (
	"time"
	"sync"
)

type CacheInterface interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
	readLoop()
}

type Cache struct {
	Entry map[string]cacheEntry
	interval time.Duration
	mutex sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	Val []byte
}

func(c *Cache)Add(key string, val []byte){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Entry[key] = cacheEntry{
		createdAt: time.Now(),
		Val: val,
	}
}

func(c *Cache)Get(key string) ([]byte, bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, exists := c.Entry[key]
	if !exists {
		return nil, false
	}
	return entry.Val, true
}

func(c *Cache)readLoop(){
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mutex.Lock()
		
		for key, entry := range c.Entry{
			if time.Since(entry.createdAt) > c.interval {
				delete(c.Entry, key)
			}
		}
		c.mutex.Unlock()
	}

}



func NewCache (interval time.Duration) *Cache{
	c := &Cache{
		Entry: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.readLoop()

	return c
}