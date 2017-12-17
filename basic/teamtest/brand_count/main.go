package main

import (
	b "./brand"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("please input the data file path")
		os.Exit(2)
	}
	start := time.Now()
	err := b.ReadAndHandle(args[len(args)-2], args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}
	elapsed := time.Now().Sub(start)
	log.Printf("total elapsed time: %f seconds", elapsed.Seconds())
	b.ListResult()
}
