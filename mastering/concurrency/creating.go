package concurrency

import (
	"fmt"
	"time"
)

func Creating() {
	for i := 0; i < 10; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println()
}
