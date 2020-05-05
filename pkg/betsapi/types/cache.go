package types

import (
	"errors"
	"github.com/robfig/cron"
	"log"
	"reflect"
	"sync"
)

// Persistent cache module
type Cache struct {
	// internal map of data
	data sync.Map

	// Cron resolution for update trigger
	UpdateInterval string

	// Reset resolution (leave nil if you don't want to clear storage)
	ResetInterval string

	// update function
	UpdateFunc func()

	// Reset function
	ResetFunc func()

	// Initialization function
	InitializeFunc func()

	// Cron handler
	c *cron.Cron
}

func (cache *Cache) LoadAs(interface{}, reflect.Type) (interface{}, bool) {
	panic("implement me")
}

// Initialize cache
// This methods runs all cronjobs and routines
func (cache *Cache) Initialize() error {
	cache.c = cron.New()

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

	//start cronjob
	cache.c.Start()

	return nil
}

// Check if entry exists in cache
func (cache *Cache) Exists(key interface{}) bool {
	_, ok := cache.data.Load(key)
	return ok
}

// Load entry from cache
func (cache *Cache) Load(key interface{}) (value interface{}, ok bool) {
	return cache.data.Load(key)
}

// Store value to cache
func (cache *Cache) Store(key interface{}, value interface{}) {
	cache.data.Store(key, value)
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

func (cache *Cache) Range(f func(key interface{}, value interface{}) bool) {
	cache.data.Range(f)
}
