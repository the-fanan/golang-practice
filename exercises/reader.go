package main

import (
	"fmt"
	"os"
	"bufio"
)
/**
	* This program keeps reading from input till the word "STOP" is entered
*/
func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for (scanner.Scan()) {
		fmt.Println(">", scanner.Text())
		if scanner.Text() == "STOP" {
			os.Exit(1)
		}
	}
}