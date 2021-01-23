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

func DoNotWaitForGoRoutinesToFinishExecution(){
	n := 10
	for i := 0; i < n; i++ {
		go func(x int) {
			fmt.Printf("No wait: %d ", x)
		}(i)
	}
}

func WaitForGoRoutinesToFinishExecution(){
	n := 10
	var waitGroup sync.WaitGroup
	for i := 0; i < n; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("wait: %d ", x)
		}(i)
	}
	waitGroup.Wait()
}