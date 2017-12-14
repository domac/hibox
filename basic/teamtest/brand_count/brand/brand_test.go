package brand

import (
	"testing"
)

func TestNullRead(t *testing.T) {
	if err := ReadAndHandle(""); err == nil {
		t.Fatalf("null test error")
	}
}
