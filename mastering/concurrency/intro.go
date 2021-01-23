package concurrency

import(
	"fmt"
	"time"
	"sync"
)

func concurrentFunction() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

func RoutinesAreUnordered(){
	go concurrentFunction()
	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	
	time.Sleep(1 * time.Second)
}

