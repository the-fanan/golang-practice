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

/**
* Order in which channels are written and read matters
 * is chan s is sent a message before chan n then chan s message must be read before chan n
*/
func undChannels(x int, s chan int, n chan int) {
	s <- x
	n <- 1
}	

func pinger(c chan string) {
	c <- "ping"
}

func printer(c chan string) {
	msg := <- c
	fmt.Println(msg)
}

func main() {
	go countDown(5, "f")//run asyncrhonously
	go countDown(6, "s")

	s := make(chan int)
	n := make(chan int)
	sum := 0

	/*

	for j := 0; j < 3; j++ {
		go undChannels(j + 1,s,n)
	}*/
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	go undChannels(5, s, n)
	go undChannels(6, s, n)
	go undChannels(9, s, n)
	//since this block reads from channels, main will wait channel to be empty before closing
	// remember, main is also a Go ROutine ;)
	for i := 0; i < 3 ; i += <- n {
		sum += <- s
	}
	fmt.Println("sum: ", sum)
}