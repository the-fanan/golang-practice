package main

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)

/**
	* This function finds the average all inputs in the command line that are numeric
*/
func main() {
	if (len(os.Args ) == 1) {
		fmt.Println("Please give one or more numbers.")
		os.Exit(1)
	}
	
	args := os.Args  
	sum := 0.0
	var err error = errors.New("All inputs must be numbers.")

	for i := 1; i < len(args); i ++ {
		n, er := strconv.ParseFloat(args[i], 64)
		if er != nil {
			fmt.Printf("%s (custom: %s)\n", err.Error(), er)
			os.Exit(1)
		}
		sum = sum + n
	}
	var k float64 = (float64) ((len(args) - 1))
	var avg float64 = sum / k
	fmt.Printf("Average of all numbers is: %f \n", avg)
}