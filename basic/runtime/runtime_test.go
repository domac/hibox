package runtime

import (
	"strconv"
	"testing"
	"time"
)

func TestRuntimeStats(t *testing.T) {

	go ReadRuntimeStats(1 * time.Second)

	go func() {

		for {
			m := make(map[string]int)
			for i := 0; i < 1024; i++ {
				m[strconv.Itoa(i)] = i
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	time.Sleep(10 * time.Second)
}
