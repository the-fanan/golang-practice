package main

import (
	"fmt"
)

func main() {
	/**
		* For Loop
	*/
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue//skip
		}
		fmt.Println(i)
		if i == 5 {
			break//break can exit out of a for loop
		}
	}

	/**
		* While loop
	*/
	j := 0
	for {
		j++
		fmt.Println(j)
		if j == 3 {
			break
		}
	}

	/**
		* Do While
	*/

	k := 0
	for l := true; l; l = k < 6 {
		fmt.Println(k)
		k++
	}

	//Range
	a := [5]int{4,8,1,6,3}
	for i,v := range a {
		fmt.Printf("index: %d, value: %d \n", i,v)
	}
}