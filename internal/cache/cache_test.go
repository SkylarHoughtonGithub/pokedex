package cache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	interval := 5 * time.Minute
	cache := NewCache(interval)

	if cache == nil {
		t.Error("NewCache should return a non-nil cache")
	}
}

func TestCacheAdd(t *testing.T) {
	interval := 5 * time.Minute
	cache := NewCache(interval)

	key := "test_key"
	value := []byte("test_value")
	cache.Add(key, value)

	retrievedValue, found := cache.Get(key)
	if !found {
		t.Errorf("Expected to find key '%s' in cache", key)
	}
	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value '%s', got '%s'", string(value), string(retrievedValue))
	}
}

func TestCacheDelete(t *testing.T) {
	interval := 5 * time.Minute
	cache := NewCache(interval)

	key := "test_key"
	value := []byte("test_value")
	cache.Add(key, value)
	cache.Delete(key)

	_, found := cache.Get(key)
	if found {
		t.Errorf("Key '%s' should have been deleted from cache", key)
	}
}

func TestCacheClear(t *testing.T) {
	interval := 5 * time.Minute
	cache := NewCache(interval)

	cache.Add("key1", []byte("value1"))
	cache.Add("key2", []byte("value2"))
	cache.Clear()

	if cache.Size() != 0 {
		t.Errorf("Cache should be empty after Clear(), but has %d items", cache.Size())
	}
}

func TestCacheSize(t *testing.T) {
	interval := 5 * time.Minute
	cache := NewCache(interval)

	if cache.Size() != 0 {
		t.Errorf("New cache should have 0 size, got %d", cache.Size())
	}

	cache.Add("key1", []byte("value1"))
	cache.Add("key2", []byte("value2"))

	if cache.Size() != 2 {
		t.Errorf("Expected cache size 2, got %d", cache.Size())
	}
}
