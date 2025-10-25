package pokeapi

import (
	"testing"
	"time"

	cache "github.com/skylarhoughtongithub/gopokedex/internal/cache"
)

func TestCommandExplore(t *testing.T) {
	// Setup a test configuration and cache
	cfg := &Config{}
	testCache := cache.NewCache(5 * time.Minute)

	// Test with no arguments
	err := CommandExplore(cfg, testCache)
	if err == nil {
		t.Error("Expected error when no location area is specified")
	}

	// Note: For more comprehensive testing of CommandExplore,
	// you'd typically use a mock HTTP client or create a test server
	// This is a minimal example and doesn't test network interactions
}

func TestCommandMap(t *testing.T) {
	// Setup a test configuration and cache
	cfg := &Config{}
	testCache := cache.NewCache(5 * time.Minute)

	// Test initial map command
	err := CommandMap(cfg, testCache)
	if err != nil {
		t.Errorf("Unexpected error in initial map command: %v", err)
	}

	// Verify that NextURL and PrevURL are set
	if cfg.NextURL == nil {
		t.Error("NextURL should be set after initial map command")
	}
}

func TestCommandMapB(t *testing.T) {
	// Setup a test configuration and cache
	cfg := &Config{
		NextURL: nil,
		PrevURL: nil,
	}
	testCache := cache.NewCache(5 * time.Minute)

	// Test map back when no previous page exists
	err := CommandMapB(cfg, testCache)
	if err != nil {
		t.Errorf("Unexpected error when no previous page exists: %v", err)
	}

	// Note: For more comprehensive testing, you'd need to simulate
	// having a previous URL and cached data
}
