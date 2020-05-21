package main

import (
	"fmt"
	"math"
)
/**
	* Write a program to print a string n times in time complexity better than O(n)
	*/
func printString(initialString string, n int) { 
	//initial the final string to be equal to the string we want to print n times
	finalString := initialString

	//here we determine what the minimum 2^k string generated should be
	minimumGpIterations := int(math.Floor(float64(n) / 2.0))

	//once finalString's length is equal to this perform the final Geometric Increase on it
	minimumStringLegnth := len(initialString) * minimumGpIterations

	//As long as the length of the final string is not greater than our minimum required length, increase it geometrically
	for len(finalString) <= minimumStringLegnth {
		//Increase the final string by geometric progrssion
		finalString = finalString + finalString
	}
	/**
		* So far what we have done will be valid for situations where 2^k evaluates to n and k >= 0
		* if n falls between 2^k and 2^(k+1) our output will be 2^k ---> You can experiment
		*/


	//Here we handle situations of when n falls between 2^k and 2^(k+1)
	
	finalStringLength := len(finalString)
	initialStringLength := len(initialString)
	expectedFinalStringLength := initialStringLength * n //n is the number of times the string is to be printed

	if (finalStringLength < expectedFinalStringLength) {	
		// We are dividing by the initial strings length to get the actual number of times left to print out
		// Remember that finalStringLength = 2^k but less than n
		arithmeticIterations := (expectedFinalStringLength/initialStringLength) - (finalStringLength/initialStringLength)
		for j := 0; j < arithmeticIterations; j++ {
			//increase final string by arithemtic progression
			finalString += initialString
		}
	}
	// the second parameter in this print statement is to verify that the string was printed the number of times specified
	fmt.Println(finalString, len(finalString)/initialStringLength)
}

func main(){
	printString("ABC", 6)
}