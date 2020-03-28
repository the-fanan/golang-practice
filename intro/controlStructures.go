package main

import (
	"fmt"
)

func main() {
	//for Loop
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		//switch
		switch i {
			case 0: fmt.Print(" Zero; ")
			case 1: fmt.Print(" One; ")
			case 2: fmt.Print(" Two; ")
			case 3: fmt.Print(" Three; ")
			case 4: fmt.Print(" Four; ")
			case 5: fmt.Print(" Five; ")
			case 6: fmt.Print(" Six; ")
			case 7: fmt.Print(" Seven; ")
			case 8: fmt.Print(" Eigth; ")
			case 9: fmt.Print(" Nine; ")
			case 10: fmt.Print(" Ten; ")
			default: fmt.Print("Unknown Number")
		}

		if i % 2 == 0 {
			fmt.Println(i, "is even")
		} else if i % 3 == 0 {
			fmt.Println(i, "is odd and divisible by 3")
		} else {
			fmt.Println(i, "is odd and not divisible by 3")
		}
	}

	//Exercise: Print numbers between 1 and 100 that are divisble by 3
	fmt.Println("Numbers from 1 to 100 divisible by 3")
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
	//Exercise: Print numbers between 1 and 100 that are divisble by 3 or 5
	fmt.Println("Fizz for 3, Buzz for 5, FizzBuzz for bot")
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
