package main

import (
	"fmt"
	"os"
)

func main() {
	/**
		* Go does not allow you include packages you would not use
		* This can be bypassed by placing an "_" before the package name
	*/
	fmt.Println("This can't compile because I included a package I did not use")
}