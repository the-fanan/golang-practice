package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	/**
		* The advantage of "os" is that it is platform independent
	*/
	// Reading user input from command line
	var f *os.File//f is a pointer variable of type os.File?
	f = os.Stdin //f now takes in os.Stdin -- still a little confused here, I thouhgt f is a pointer??
	defer f.Close()//close f just before the program terminates
	scanner := bufio.NewScanner(f)// pass Stin instance to scanner to scnan
	for (scanner.Scan()) {//while snner is active, read input
		fmt.Println(">", scanner.Text())//print entered text
	}
}