package cache

import (
	"log"

	lru "github.com/shriramsharma/cache/lru"
)

// Cache interface
type Cache interface {
	Put(key interface{}, value interface{}) (bool, error)
	Get(key interface{}) (interface{}, error)
	Purge() (bool, error)
	Delete(key interface{}) (bool, error)
}

// New instantiates cache of a defined size
func New(size int) (Cache, error) {
	cache, err := lru.NewLRU(size)
	if err != nil {
		log.Fatal(err)
	}
	return cache, nil
}
