package main

import (
	"fmt"
)

func main() {
	Debug("it's expensive")
	if Dev {
		fmt.Println("we are in develop mode")
	}
}
