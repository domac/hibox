package main

import (
	"regexp"
)

func main() {

	matched, _ := regexp.MatchString(`\d+`, "123456")
	println(matched)

	m2, _ := regexp.MatchString(`\p{N}`, "123")
	println(m2)

}
