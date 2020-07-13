package main

import (
	"fmt"
)

func printString(s string, n int) string { 
	fs := s
	mGI := n/2
	mSL := len(s) * mGI

	for len(fs) <= mSL {
		fs = fs + fs
	}

	fsL := len(fs)
	sL := len(s)
	eFSL := sL * n

	if (fsL < eFSL) {	
		newN := (eFSL/sL) - (fsL/sL)
		fs += printString(s, newN)
	}
	return fs
}

func main(){
	fmt.Println(printString("ABC", 3))
}