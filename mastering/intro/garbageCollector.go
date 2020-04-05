package main

import (
	"fmt"
	"time"
	"runtime"
)

func printStatistics(m runtime.MemStats) {
	runtime.ReadMemStats(&m)
	fmt.Println("mem.Alloc:", m.Alloc)
	fmt.Println("mem.TotalAlloc:", m.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", m.HeapAlloc)
	fmt.Println("mem.NumGC:", m.NumGC)
	fmt.Println("-----")
}

func main() {
	var m runtime.MemStats
	printStatistics(m)
	//trigger garbage collector by creating multiple huge Go slices
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStatistics(m)
	//allocate more memory again
	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	
	printStatistics(m)
}