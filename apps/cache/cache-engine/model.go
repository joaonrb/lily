//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
// This Cache system is done for test purposes mostly
//
package cengine

import (
	"time"
	"sync"
)

const (
	DEFAULT_DURATION = time.Hour * 24
	MAX_DURATION     = time.Hour * 24 * 30
)

var (
	writeLock  = sync.RWMutex{}
	readCache  = map[string]*cacheEntry{}
	writeCache = readCache
)

type cacheEntry struct {
	data            interface{}
	expirationDate  time.Time
}

func Put(key string, value interface{}, expiration... time.Duration) {
	writeLock.RLock()
	entry, ok := readCache[key]
	writeLock.RUnlock()
	if !ok {
		entry = &cacheEntry{}
	}

	var expirationDate time.Time
	if len(expiration) > 0 {
		if expiration[0] > MAX_DURATION {
			expirationDate = time.Now().UTC().Add(MAX_DURATION)
		} else {
			expirationDate = time.Now().UTC().Add(expiration[0])
		}
	} else {
		expirationDate = time.Now().UTC().Add(DEFAULT_DURATION)
	}
	entry.data = value
	entry.expirationDate = expirationDate
	writeLock.Lock()
	writeCache[key] = entry
}

func Get(key string) (interface{}, bool) {
	writeLock.RLock(); defer writeLock.Unlock()
	return readCache[key]
}