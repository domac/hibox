package tw

import (
	"testing"
	"time"
)

func Test_TimingWheel(t *testing.T) {
	w := NewSimpleClock(1*time.Second, 60)
	defer w.Stop()

	<-w.After(5 * time.Second)

	var err interface{}
	func() {
		defer func() {
			err = recover()
		}()
		w.After(1 * time.Second)
	}()

	if err != nil {
		t.Fail()
	}
}
