package concurrency

import (
	"fmt"
	"math/rand"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func First(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- Random(min, max)
	}
}

func Second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func Third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}
