package concurrency

import (
	"fmt"
)

func WriteToChannel1(c chan int, x int) {
	fmt.Println(x)
	c <- x //blocks rest of function as nobody reads what is written here
	close(c)
	fmt.Println(x) //never gets executed cos program gets terminated before line c <- x gets unblocked
}

func WriteToChannel2(c chan int, x int) {
	fmt.Println("1", x)
	c <- x //blocks rest of function as nobody reads what is written here
	close(c)
	fmt.Println("2", x) //never gets executed cos program gets terminated before line c <- x gets unblocked
}
