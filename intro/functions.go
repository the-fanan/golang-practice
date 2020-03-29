package main

import (
	"fmt"
)
//func NAME (...ARGS) RETURN_TYPE {}
func average(x[] float64) float64 {
	total := 0.0
	for _, value := range x {
		total += value
	}

	return total/float64(len(x))
}
//Go functions can return multiple values
func multReturn() (int, int){
	return 2,5
}
//variadic functions (multiple parameters)
func summer(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}

	return total
}
//closure
func outerFunction(x int) func() int {
	return func() (ret int) {
		x++
		return x
	}
}
//recursion
func countDown(x int) {
	fmt.Println(x)
	if (x == 0) {
		return
	} else {
		countDown(x - 1)
	}
}
//using pointers 
func pointerHandler(xP *int, yP *int) int {
	*xP = 5
	*yP = 10

	return *xP + *yP
}

func main() {
	x := []float64{98,93,77,82,83}
	fmt.Println(average(x))

	t,f := multReturn()
	fmt.Println(t,f)
	fmt.Println(summer(1,2,3,4,5,6))
	//passing a slice
	xs := []int{11,23,32}
	fmt.Println(summer(xs...))
	sc := outerFunction(5)
	sc()//6
	sc()//7
	fmt.Println(sc())

	countDown(5)
	pP := 0
	mP := new(int)
	fmt.Println(pointerHandler(&pP, mP))
	//using defer to recover from runtime error
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}