package runtime

import (
	"fmt"
	r "runtime"
	"time"
)

func metric(startInfo, endInfo *r.MemStats) (string, error) {
	res := fmt.Sprintf("goroutines :%d| current heap objects :%d| current malloc :%d", r.NumGoroutine(), endInfo.HeapObjects, endInfo.Mallocs)
	return res, nil
}

//定时统计运行时性能信息
func ReadRuntimeStats(d time.Duration) {
	var startM r.MemStats
	var currentM r.MemStats
	r.ReadMemStats(&startM)
	for {
		time.Sleep(d)
		r.ReadMemStats(&currentM)
		res, err := metric(&startM, &currentM)
		if err != nil {
			return
		}
		fmt.Println(res)
		startM = currentM
	}
}
