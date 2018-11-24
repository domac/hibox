package main

import (
	"regexp"
)

func main() {
	m1, _ := regexp.MatchString(`^1`, "1\n2\n")
	println(m1)

	m2, _ := regexp.MatchString(`^2`, "1\n2\n")
	println(m2)

	m3, _ := regexp.MatchString(`(?m)^2`, "1\n2\n")
	println(m3)

	m4, _ := regexp.MatchString(`1$`, "1\n")
	println(m4)

	m5, _ := regexp.MatchString(`(?m)1$`, "1\n")
	println(m5)
}
