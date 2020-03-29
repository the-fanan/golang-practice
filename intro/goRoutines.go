package main

import (
	"fmt"
)

func countDown(x int, h string) {
	fmt.Printf("%s: %d\n", h, x)
	if x == 0 {
		return
	}
	countDown(x - 1, h)
}

func main() {
	go countDown(5, "f")//run asyncrhonously
	go countDown(6, "s")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}