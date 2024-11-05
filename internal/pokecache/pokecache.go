// Package pokecache - cache for the pokeapi
package pokecache

import (
	"sync"
	"time"
)

// Pokecache -
type Pokecache struct {
	entries map[string]CacheEntry
	mux     *sync.RWMutex
}

// CacheEntry -
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Add -
func (p *Pokecache) Add(key string, val []byte) {
	p.mux.Lock()
	p.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	p.mux.Unlock()
}

// Get -
func (p *Pokecache) Get(key string) ([]byte, bool) {
	p.mux.RLock()
	cacheEntry, exists := p.entries[key]
	p.mux.RUnlock()
	if !exists {
		return nil, false
	}

	return cacheEntry.val, true
}

// reapLoop -
func (p *Pokecache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		now := time.Now()
		p.reap(now, interval)
	}
}

func (p *Pokecache) reap(now time.Time, interval time.Duration) {
	for k, v := range p.entries {
		if now.Sub(v.createdAt) > interval {
			p.mux.Lock()
			delete(p.entries, k)
			p.mux.Unlock()
		}
	}
}

// NewCache -
func NewCache(interval time.Duration) *Pokecache {
	mux := &sync.RWMutex{}

	p := &Pokecache{
		entries: make(map[string]CacheEntry),
		mux:     mux,
	}

	go p.reapLoop(interval)

	return p
}
