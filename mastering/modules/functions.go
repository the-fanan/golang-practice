package main

import (
	"fmt"
)

//named return values
func returnNamed(x, y int) (str string, min, max int) {
	if x < y {
		min = x
		max = y
		str = fmt.Sprintf("x(%d) is less than y(%d)", x, y)
	} else {
		min = y
		max = x
		str = fmt.Sprintf("y(%d) is less than x(%d)", y, x)
	}

	return
	//if the named variables are not appended to the return statement, they are returned in the order in whch they are defined.
	//if they are appened, they are returned in the roder in which they are appended
}

func main() {
	str,min,max := returnNamed(5,6)
	fmt.Println(min,max,str)
}
