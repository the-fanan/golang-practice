package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "89sdsd8989aaae"
	r := regexp.MustCompile("[0-9][A-Za-z]")
	if r.MatchString(s) {
		fmt.Println(r.FindAllString(s,len(s)))
	} else {
		fmt.Println("No alphabets found in strings")
	}
} 
/**
 "a" - match a
 * - match 0 or more
 + - match 1 or more
 ? - match 0 or 1
*/