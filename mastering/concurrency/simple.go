package concurrency

import (
	"fmt"
)

func FirstConcurrent() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}
