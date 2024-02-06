package add

import "testing"

func TestAdd(t *testing.T) {
	if expected, got := 4, add(2, 2); expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
