package main

import (
	"io"
	"os"
)

func main() {
	// := is the "short assignment operatpr". It cannot be used outside a function. When used, a value must be declared
	// var can be used outside a function
	// conventionally, var is used for declaring global variables
	s := ""
	args := os.Args
	if (len(args) == 1) {
		s = "Give me an argument!"
	} else {
		s = args[1]
	}

	io.WriteString(os.Stdout, s)
	io.WriteString(os.Stdout, "\n")
	/*
		* io.WriteString(DESTINATION_FILE, CONTENT)
		* Remember that unix treats everything as a file (printers, computers, etc.)
	*/
}