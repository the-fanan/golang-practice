package main

import (
	"fmt"
)
func main() {
	/**
		
	* Integers
	
	var unsigned8 uint8 // same as byte
	var unsigned16 uint16
	var unsigne32 uint32
	var unsigned64 uint64
	var signed8 int8
	var signed16 int16
	var signed32 int32 //same as rune
	var signed64 int64
	var general int


	* Floats
	
	var floating32 float32
	var floating64 float64 

	NaN for 0/0
	*/
	//arithmetic operations in Go
	var a int = 10
	var b int = 2
	var c int = 3
	var d float32 = float32(a) / float32(c)
	fmt.Printf("Subtraction: %d - % d = %d \n", b, c, c - b)
	fmt.Printf("Addition: %d + %d = %d \n", a, c, a + c)
	fmt.Printf("Division: %d / %d = %f \n", a, c, d)
	fmt.Printf("Multiplication: %d * %d = %d \n", a, c, a * c)
	fmt.Printf("Modulo: %d %% %d = %d \n", a, b, a % b)
	fmt.Printf("Modulo: %d %% %d = %d \n", a, c, a % c)

	// Strings
	var s string = "Hello World!"
	fmt.Println(s)
	fmt.Println(len(s))

	// Bolleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	//constants - value cannot be changed later
	const con int = 100
	fmt.Println("A constant", con)	
}