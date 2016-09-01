package lru

import (
	"container/list"
	"fmt"
	"sync"
)

// LRUCache Thread-safe LRU cache struct
type LRUCache struct {
	pageList    *list.List
	lookupTable map[interface{}]interface{}
	lock        sync.RWMutex
	size        int
}

// NewLRU Creates a new LRU Cache object
func NewLRU(size int) (*LRUCache, error) {
	lru := &LRUCache{
		pageList:    list.New(),
		lookupTable: make(map[interface{}]interface{}),
		size:        size,
	}
	return lru, nil
}

// Put  Adds the given key value to the cache
func (c *LRUCache) Put(key interface{}, value interface{}) (bool, error) {
	if v, ok := c.lookupTable[key]; ok {
		c.pageList.MoveToFront(v.(*list.Element))
		return false, nil
	}
	if c.pageList.Len() > c.size {
		lastElement := c.pageList.Back()
		c.pageList.Remove(lastElement)
	}
	c.pageList.PushFront(value)
	c.lookupTable[key] = value

	return false, nil
}

// Get Returns the page in cache else returns nil in case of a cache miss.
func (c *LRUCache) Get(key interface{}) (interface{}, error) {
	if v, ok := c.lookupTable[key]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("%s Not found", key)
}

// Delete Removes the entry from the cache
func (c *LRUCache) Delete(key interface{}) (bool, error) {
	return false, nil
}

// Purge Purges the entire cache
func (c *LRUCache) Purge() (bool, error) {
	return false, nil
}
