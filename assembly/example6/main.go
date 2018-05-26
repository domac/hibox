package main

import (
	"github.com/domac/hibox/assembly/example6/gid"
	"log"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			id := gid.Get()
			log.Println(id)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
