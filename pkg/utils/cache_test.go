package utils

import (
	"strconv"
	"testing"
	"time"
)

func TestCache_Init(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCache_Store(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second * 1})
	if !cache.Exists("test_cache") {
		t.Fatal("stored item doesnt exist")
	}

	time.Sleep(time.Second * 1)
	if cache.Exists("test_cache") {
		t.Fatal("item should not exist as ttl expired")
	}
}

func TestCache_Load(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second * 1})
	if !cache.Exists("test_cache") {
		t.Fatal("stored item doesnt exist")
	}

	if val, found := cache.Load("test_cache"); found {
		if val.(string) != "test" {
			t.Fatal("unexpected value for stored key")
		}
	} else {
		t.Fatal("stored value not found in cache")
	}
}

func TestCache_Exists(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second * 1})
	if !cache.Exists("test_cache") {
		t.Fatal("stored item doesnt exist")
	}
}

func TestCache_Delete(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second * 10})
	if !cache.Exists("test_cache") {
		t.Fatal("stored item doesnt exist")
	}

	cache.Delete("test_cache")
	if cache.Exists("test_cache") {
		t.Fatal("stored item wasn't deleted")
	}
}

func TestCache_DeleteExpiredItems(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second})
	if !cache.Exists("test_cache") {
		t.Fatal("stored item doesnt exist")
	}

	time.Sleep(time.Second)
	cache.DeleteExpiredItems()
	if cache.Exists("test_cache") {
		t.Fatal("stored item should not exist")
	}
}

func TestCache_Range(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++{
		seq := strconv.FormatInt(int64(i), 10)
		cache.Store("test"+seq, "val"+ seq)
	}

	cache.Range(func(key interface{}, value interface{}) bool {
		newVal := value.(string) + "_new"
		cache.Store(key, newVal)
		return true
	})

	for i := 0; i < 5; i++{
		seq := strconv.FormatInt(int64(i), 10)
		if val, found := cache.Load("test"+seq); found{
			if val.(string) != "val"+seq+"_new"{
				t.Fatal("item doesnt affected by range function")
			}
		} else{
			t.Fatal("missing item")
		}
	}
}

func TestCache_Count(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second})
	if cache.Count() != 1 {
		t.Fatal("wrong number of keys count")
	}
}

func TestCache_Keys(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second})
	if cache.Count() != 1 {
		t.Fatal("wrong number of keys count")
	}
}

func TestCache_Clear(t *testing.T) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatal(err)
	}

	cache.Store("test_cache", "test", &StoreOption{Ttl: time.Second})
	if !cache.Exists("test_cache") {
		t.Fatal("stored item doesnt exist")
	}

	time.Sleep(time.Second)
	cache.Clear()
	if cache.Exists("test_cache") {
		t.Fatal("stored item should not exist")
	}
}

func TestGetLocalCache(t *testing.T) {
	cache := GetLocalCache()
	if cache == nil {
		t.Fatal("local cache is nil")
	}
}
