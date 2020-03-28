package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	// converting input from command line
	if (len(os.Args) == 1) {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)
	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	/**
		* strconv.ParseFloat() returns two value. The first is the command line, the second is an error variable
		* Using an underscore means we want to discard a value. Underscore is called BLANK IDENTIFIER
	*/
}