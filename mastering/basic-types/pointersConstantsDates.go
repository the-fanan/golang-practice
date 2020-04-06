package main

import (
	"fmt"
)
/**
	*type VARIABLE_NAME DATA_TYPE -- is a way of defining a new named type that uses the same
underlying type as an existing one. This is mainly used for differentiating
between different types that might use the same kind of data.
*/
type Digit int
//constants
const Pi = 3.1415926
const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
	)

func main() {
	fmt.Println(Pi)
	//constant generator 'iota'
	/**
		* This code below is equivalent to:
		const (
		zero = 0
		one = 1
		two = 2
		three = 3
	)
	*/
	const (
		zero Digit = iota
		one
		two
		three
	)
	fmt.Println(two)

	/**
		* Pointers
	*/
	v := 7
	n := new(int)//create a pointer to an int
	n = &v
	t := &v //t is assigned a pointer to v
	fmt.Println(t, n)//notice that t and n are memory addresses
	fmt.Println(*t, *n)//* -- access the actual value the pointer points to
	*n = 9 // change the value of what n points to
	fmt.Println(v,*n,*t)
}