package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteToChannel(t *testing.T) {
	c := make(chan int)
	go WriteToChannel1(c, 10)
	time.Sleep(1 * time.Second)
}

func TestReadChannel(t *testing.T) {
	c := make(chan int)
	go WriteToChannel2(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
