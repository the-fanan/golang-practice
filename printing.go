package main

import (
	"fmt"
)

func main() {
	fmt.Print("No new line")
	fmt.Println("New line added")
	fmt.Printf("A verb %d \n", 10)

	v1 := "123"
	v2 := 123
	v3 := "Have a nice day"
	v4 := "abc"

	fmt.Println(v1, v2, v3, v4)
	fmt.Printf("%s %d '%s' %s\n", v1, v2, v3, v4)
}