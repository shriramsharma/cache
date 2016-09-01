package lru

import "testing"

func TestNewLRU(t *testing.T) {
	if _, err := NewLRU(0); err == nil {
		t.Errorf("Expected an invalid size error")
	}
	size := 128
	cache, err := NewLRU(size)
	if err != nil {
		t.Errorf("%s", err)
	}
	if cache == nil {
		t.Errorf("Error in instantiating cache")
	}
	if cache.pageList == nil {
		t.Errorf("pageList is nil")
	}
	if cache.lookupTable == nil {
		t.Errorf("lookupTable is nil")
	}
	if cache.size != size {
		t.Errorf("Incorrect cache size")
	}

}
