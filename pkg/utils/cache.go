package utils

import (
	"errors"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Persistent cache module
type Cache struct {
	data           sync.Map
	ttl            time.Duration
	UpdateInterval string
	ResetInterval  string
	UpdateFunc     func()
	ResetFunc      func()
	InitializeFunc func()
	c              *cron.Cron
}

// Init cache
func (cache *Cache) Init() error {
	cache.c = cron.New(cron.WithSeconds())
	if cache.ttl == 0 {
		cache.ttl = time.Hour
	}

	if cache.UpdateInterval != "" {
		if cache.UpdateFunc == nil {
			return errors.New("update function is not defined")
		}
		if _, err := cache.c.AddFunc(cache.UpdateInterval, cache.UpdateFunc); err != nil {
			log.Fatal(err)
		}
	}

	if cache.ResetInterval != "" {
		if cache.ResetFunc == nil {
			return errors.New("reset function is not defined")
		}
		if _, err := cache.c.AddFunc(cache.ResetInterval, cache.ResetFunc); err != nil {
			log.Fatal(err)
		}
	}

	if cache.InitializeFunc != nil {
		cache.InitializeFunc()
	}

	cache.c.Start()

	return nil
}

// Check if entry exists in cache
func (cache *Cache) Exists(key interface{}) bool {
	storedItem, ok := cache.data.Load(key)
	if !ok {
		return false
	}

	now := int32(time.Now().Unix())
	if asserted, ok := storedItem.(*cacheItem); ok && asserted.validUntil > now {
		return true
	}
	cache.data.Delete(key)
	return false
}

// Load entry from cache
func (cache *Cache) Load(key interface{}) (value interface{}, ok bool) {
	storedItem, ok := cache.data.Load(key)
	if !ok {
		return nil, false
	}

	now := int32(time.Now().Unix())
	if asserted, ok := storedItem.(*cacheItem); ok && asserted.validUntil > now {
		return asserted.value, true
	}

	return nil, false
}

// Store value to cache
func (cache *Cache) Store(key interface{}, value interface{}, options ...*StoreOption) {
	if key == "" || value == nil {
		return
	}

	var storeOption *StoreOption
	if len(options) > 0 && options[0].Ttl > -1 {
		storeOption = options[0]
	} else {
		storeOption = &StoreOption{Ttl: cache.ttl}
	}

	log.Debugf("Storing %s for %d", key, storeOption.Ttl/1000000000)

	cacheItem := newCacheItem(value, storeOption)
	cache.data.Store(key, cacheItem)
}

// Delete entry from cache
func (cache *Cache) Delete(key interface{}) {
	cache.data.Delete(key)
}

// Clear all entries in cache
func (cache *Cache) Clear() {
	cache.data.Range(func(key interface{}, _ interface{}) bool {
		cache.data.Delete(key)
		return true
	})
}

// Return number of entries in cache
func (cache *Cache) Count() int {
	length := 0
	cache.data.Range(func(_ interface{}, _ interface{}) bool {
		length++
		return true
	})

	return length
}

// Return all keys in cache
func (cache *Cache) Keys() []interface{} {
	keys := make([]interface{}, 0)
	cache.data.Range(func(key interface{}, _ interface{}) bool {
		keys = append(keys, key)
		return true
	})

	return keys
}

func (cache *Cache) Range(f func(key interface{}, value interface{}) bool) {
	cache.data.Range(func(key, value interface{}) bool {
		v := value.(*cacheItem).value
		return f(key, v)
	})
}

func (cache *Cache) DeleteExpiredItems() {
	cache.data.Range(func(key interface{}, _ interface{}) bool {
		cache.Exists(key) // will delete if ttl has been reached
		return true
	})
}

// A StoreOption sets options such as ttl ..
type StoreOption struct {
	// classic time.Duration - ie. 5 * time.Minute
	Ttl time.Duration
	// Store calls never checks if the key does exist - it is up to wrapping logic to determine
	// if Store should be called. This value helps to determine the outcome
	ForceReload bool
	// this value gelps to determine the correct key - explicit choice
	OverrideKey string
}

type cacheItem struct {
	value      interface{}
	validUntil int32
}

func newCacheItem(value interface{}, storeOptions *StoreOption) *cacheItem {
	return &cacheItem{
		value,
		int32(time.Now().Add(storeOptions.Ttl).Unix()),
	}
}

var localCacheInstance *Cache
var getLocalCacheOnce sync.Once

// Returns intialized local cache instance
func GetLocalCache() *Cache {
	getLocalCacheOnce.Do(func() {
		localCacheInstance = &Cache{}
		err := localCacheInstance.Init()
		if err != nil {
			log.Fatal(err)
		}
	})
	return localCacheInstance
}
