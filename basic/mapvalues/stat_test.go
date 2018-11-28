package mapvalues

import (
	"testing"
	"time"
)

func TestStat(t *testing.T) {

	routes := make(map[uint32]int)

	for i := 0; i < 100; i++ {
		routes[uint32(i)] = i
	}

	stat, _ := NewConnscStat(routes)
	stat.TempStart()

	for i := 0; i < 10; i++ {
		stat.UpdateAllowCount(uint32(i % 2))
		stat.UpdateDenyCount(uint32(i % 3))
		time.Sleep(time.Millisecond * 1)
	}

	stat.PrintStats()

	time.Sleep(10 * time.Second)
}

func BenchmarkStat(b *testing.B) {
	b.ReportAllocs()

	routes := make(map[uint32]int)

	for i := 0; i < 100; i++ {
		routes[uint32(i)] = i
	}

	stat, _ := NewConnscStat(routes)
	stat.TempStart()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < b.N; i++ {
				stat.UpdateAllowCount(uint32(i % 2))
				stat.UpdateDenyCount(uint32(i % 3))
				time.Sleep(time.Millisecond * 1)

				//stat.PrintStats()
			}
		}
	})
}
