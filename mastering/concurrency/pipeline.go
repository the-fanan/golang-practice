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

/**
unbufferred channels were used for this pipeline.
Unbuffered channels are used for synchronous communication while buffered channelsare used for asynchronous communication.
An unbuffered channel will block function from continuing execution until it is read or written to. A buffered channel does not block function from continuing execution.

Both types of channels will however prevent a function from closing until they are closed
*/
