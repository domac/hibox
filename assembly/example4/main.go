package main

import (
	"fmt"
	"github.com/domac/hibox/assembly/example4/add"
)

func main() {
	r := add.AddNum(32, 256)
	fmt.Println("result is :", r)
}
