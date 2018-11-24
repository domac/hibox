package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Printf("%s\n", regexp.MustCompile(`\d+`).FindAllString("12 34 56", -1))
	fmt.Printf("%s\n", regexp.MustCompile(`\d+`).FindAllString("12 34 56", 0))
	fmt.Printf("%s\n", regexp.MustCompile(`\d+`).FindAllString("12 34 56", 1))
	fmt.Printf("%s\n", regexp.MustCompile(`\d+`).FindAllString("12 34 56", 2))
	fmt.Printf("%s\n", regexp.MustCompile(`\d+`).FindAllString("12 34 56", 3))
	fmt.Printf("%s\n", regexp.MustCompile(`\d+`).FindAllString("12 34 56", 4))
}
