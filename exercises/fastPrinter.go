package main

import (
	"fmt"
)

func printString(s string, n int) string { 
	if n == 0 {
		return ""
	}
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
	word := "A"
	repetitions := 4
	if printString(word, repetitions) == repeatString(word, repetitions) {
		fmt.Println("TRUE")
		fmt.Println(repeatString(word, repetitions))
	} else {
		fmt.Println("FALSE")
		fmt.Println(printString(word, repetitions))
		fmt.Println(repeatString(word, repetitions))
	}
}

//repeating function with different approach
func repeatString(s string, n int) string {
	if n == 1 {
		return s
	}

	running := s
	for true {
		running = running + running
		if len(running) * 2 >= len(s) * n  {
			break
		}
	}

	if len(running) == len(s) * n {
		return running
	}

	running = running + repeatString(s, n - (len(running)/len(s)))
	return running
}