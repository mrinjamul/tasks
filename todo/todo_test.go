package todo

import (
	"testing"
)

// TestGetHomeDir tests
func TestGetHomeDir(t *testing.T) {
	out := GetHomeDir()
	if out == "" || len(out) == 0 {
		t.Errorf("Want strings but got nil")
	}
}
