package main

import (
	"fmt"
	"regexp"
)

//分组演示

func main() {
	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	fmt.Println(re.ReplaceAllString("2018-12-24", "$2/$3/$1"))

	myExp := regexp.MustCompile(`(?P<first>\d+).(\d+).(?P<third>\d+)`)
	match := myExp.FindStringSubmatch("1234.5678.9")

	fmt.Printf("%s\n", match)

	result := make(map[string]string)
	for i, name := range myExp.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	fmt.Printf("by name : %s %s \n", result["first"], result["third"])

}
