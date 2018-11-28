package mapvalues

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrRountMapNull = errors.New("router map is null")
	statTime        = time.Second //分钟间隔
)

type ConnscStat struct {
	AllowCalls  map[uint32]uint64
	RejectCalls map[uint32]uint64
	cmdcount    int

	MissCalls map[uint32]uint64
	mu        sync.RWMutex

	allowMutex  sync.RWMutex
	rejectMutex sync.RWMutex
	missMutex   sync.RWMutex

	stop bool
}

func NewConnscStat(routeMap map[uint32]int) (*ConnscStat, error) {

	if routeMap == nil || len(routeMap) == 0 {
		return nil, ErrRountMapNull
	}

	cmdcount := len(routeMap)

	log.Printf("---cmdcount:%d", cmdcount)

	cs := &ConnscStat{
		AllowCalls:  make(map[uint32]uint64, cmdcount),
		RejectCalls: make(map[uint32]uint64, cmdcount),
		MissCalls:   make(map[uint32]uint64),
		stop:        true,
		cmdcount:    cmdcount,
	}

	for cmd, _ := range routeMap {
		cs.AllowCalls[cmd] = 0
		cs.RejectCalls[cmd] = 0
	}

	return cs, nil
}

func (cs *ConnscStat) PrintStats() {
	cs.mu.Lock()
	for cmd, _ := range cs.AllowCalls {
		log.Printf(">>> cmd [%d] allow  calls count: %d\n", cmd, cs.AllowCalls[cmd])
		cs.AllowCalls[cmd] = 0
	}

	for cmd, _ := range cs.RejectCalls {
		log.Printf(">>> cmd [%d] reject calls count: %d\n", cmd, cs.RejectCalls[cmd])
		cs.RejectCalls[cmd] = 0
	}

	for cmd, _ := range cs.MissCalls {
		log.Printf(">>> cmd [%d] miss   calls count: %d\n", cmd, cs.MissCalls[cmd])
		cs.MissCalls[cmd] = 0
	}
	cs.MissCalls = make(map[uint32]uint64)
	cs.mu.Unlock()

}

func (cs *ConnscStat) TempStart() {
	cs.stop = false
}

func (cs *ConnscStat) Start() {

	cs.stop = false

	log.Printf(">>>> connsc stat cmd count :%d\n", cs.cmdcount)

	ticker := time.NewTicker(statTime)
	for {
		select {
		case <-ticker.C:
			//统计日志输出
			cs.PrintStats()
		}
	}
}

func (cs *ConnscStat) UpdateAllowCount(cmd uint32) {
	if cs.stop {
		return
	}
	cs.allowMutex.Lock()
	currentCount := cs.AllowCalls[cmd]
	cs.AllowCalls[cmd] = atomic.AddUint64(&currentCount, 1)
	cs.allowMutex.Unlock()
}

func (cs *ConnscStat) UpdateDenyCount(cmd uint32) {
	if cs.stop {
		return
	}
	cs.rejectMutex.Lock()
	currentCount := cs.RejectCalls[cmd]
	cs.RejectCalls[cmd] = atomic.AddUint64(&currentCount, 1)
	cs.rejectMutex.Unlock()

}

func (cs *ConnscStat) UpdateMissCount(cmd uint32) {
	if cs.stop {
		return
	}
	currentCount, ok := cs.MissCalls[cmd]
	if !ok {
		currentCount = 0
	}
	cs.missMutex.Lock()
	cs.MissCalls[cmd] = atomic.AddUint64(&currentCount, 1)
	cs.missMutex.Unlock()
}
