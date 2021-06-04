package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	go FirstConcurrent()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(1 * time.Second)
}
