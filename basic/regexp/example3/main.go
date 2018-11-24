package main

import (
	"regexp"
)

func main() {
	m1, _ := regexp.MatchString(`abc`, "aBc")
	println(m1)

	m2, _ := regexp.MatchString(`a(?i)bc`, "aBc")
	println(m2)

	m3, _ := regexp.MatchString(`a(?i)b(?i)c`, "aBC")
	println(m3)

	m4, _ := regexp.MatchString(`a(?i)b(?-i)c`, "aBC")
	println(m4)
}
