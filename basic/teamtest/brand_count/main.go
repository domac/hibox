package main

import (
	b "./brand"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("please input the data file path")
		os.Exit(2)
	}
	b.ReadAndHandle(args[1])
}
