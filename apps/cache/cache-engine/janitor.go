//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package cengine

import (
	"sync"
	"time"
)

var busy = sync.Mutex{}

func init() {
	ticker := time.NewTimer(time.Second*30)
	go func() {
		for _ := range ticker.C {
			busy.Lock()
			runClean(time.Now().UTC())
			busy.Unlock()
		}
	}()
}

func runClean(now time.Time) {
	writeCache = map[string]*cacheEntry{}
	for key, entry := range readCache {
		if now.After(entry.expirationDate) {
			writeLock.Lock()
			writeCache[key] = entry
		}
	}
	readCache = writeCache
}
