package main

import (
	"io"
	"os"
)

func main() {
	s := ""
	args := os.Args
	if len(args) == 1 {
		s = "Please give me one argument!"
	} else {
		s = args[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, s)
	io.WriteString(os.Stderr, "\n")
	/**
		* standard input is 0
		* standard output is 1
		* standard error is 2
	*/
}