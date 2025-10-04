package mast

import (
	"fmt"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	cacheDir      = "./cache/mast"
	cacheDuration = 2 * time.Hour
)

type Storage struct {
	cache *cache.Cache
}

func NewStorage() *Storage {
	return &Storage{
		cache: cache.New(cacheDuration, 10*time.Minute),
	}
}

func (s *Storage) Start() error {
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	files, err := os.ReadDir(cacheDir)
	if err != nil {
		return fmt.Errorf("could not read cache directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		info, err := file.Info()
		if err != nil {
			continue
		}

		modTime := info.ModTime()
		elapsed := time.Since(modTime)

		if elapsed < cacheDuration {
			remaining := cacheDuration - elapsed
			s.cache.Set(file.Name(), true, remaining)
		}
	}

	return nil
}

func (s *Storage) Store(key string, value any, duration time.Duration) {
	s.cache.Set(key, value, duration)
}

func (s *Storage) Get(key string) bool {
	_, exists := s.cache.Get(key)
	return exists
}
