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
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)// pass Stin instance to scanner to scnan
	for (scanner.Scan()) {//while snner is active, read input
		fmt.Println(">", scanner.Text())//print entered text
	}
}