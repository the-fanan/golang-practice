package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		// GO allows you to create your own type of errors
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func main() {
	var b int
	b = 2
	err := returnError(1, b)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}
	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}
	if err.Error()/* .Error() converts error variable to string*/ == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}