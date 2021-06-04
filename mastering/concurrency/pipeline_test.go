package concurrency

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	A := make(chan int)
	B := make(chan int)

	go First(4, 55, A) //A acts as write only here
	go Second(B, A)    //A acts as read while B acts as write
	Third(B)           //B acts as read here
}
