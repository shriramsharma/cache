package cache

import "testing"

func TestNew(t *testing.T) {
	if _, err := New(-200); err == nil {
		t.Errorf("Expected and invalid size error")
	}
	if _, err := New(0); err == nil {
		t.Errorf("Expected and invalid size error wih size passed as 0")
	}
	_, err := New(128)
	if err != nil {
		t.Errorf("%s", err)
	}
}
