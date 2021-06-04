package concurrency

import (
	"fmt"
	"sync"
)

func Waiting() {
	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", &waitGroup)
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("%#v\n", &waitGroup)
	waitGroup.Wait()
	//This allows the program to wait for routines to be completed but still does not ensure order

	/**
	* CONCURRENCY ERRORS
	*
	* Race: occurs when concurrent processes are sharing a single resource without any mechanism to control the access
	*
	* Deadlock: occurs when concurrent processes share a resource and can lock the resource but two processes accidentally lock a resource at the same time preventing either of them from accessinf the resource

	 */
}
