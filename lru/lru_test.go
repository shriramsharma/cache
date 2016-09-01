package lru

import "testing"

func TestNewLRU(t *testing.T) {
	if _, err := NewLRU(0); err == nil {
		t.Errorf("Expected an invalid size error")
	}

}
