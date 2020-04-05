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
	var f *os.File//f is a pointer variable of struct os.File
	//os.Stdin is an open file pointing to Standard input. It is basically a pointer so next line is valid
	/*
	var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	)

	func NewFile(fd uintptr, name string) *File ---- returns pointer to struct File
	*/
	f = os.Stdin
	defer f.Close()//close f just before the program terminates
	scanner := bufio.NewScanner(f)// pass Stin instance to scanner to scnan
	for (scanner.Scan()) {//while snner is active, read input
		fmt.Println(">", scanner.Text())//print entered text
	}
}