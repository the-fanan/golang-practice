package main

import(
	"fmt"
)
//tw
/**
	* This file is just for me to try out stuff
*/
func mergeSort(a *[]int, ta *[]int, l int, r int) int{
	if l == r {
		return l
	} else {
		m := int((l + r)/2)
		l := mergeSort(a, ta, l, m)
		r := mergeSort(a, ta, m+1, r)
		merger(a, ta, l, m + 1, r + 1)
		return l
	}
}

func merger(a *[]int, ta *[]int, l int, r int, e int) {
	t := l 
	left := l
	right := r
	for left < r && right < e {
		if (*a)[left] < (*a)[right] {
			(*ta)[t] = (*a)[left]
			left++
		} else {
			(*ta)[t] = (*a)[right]
			right++
		}
		t++
	}

	for left < r {
		(*ta)[t] = (*a)[left]
		left++
		t++
	}

	for right < e {
		(*ta)[t] = (*a)[right]
		right++
		t++
	}

	for i := l; i < e; i++ {
		(*a)[i] = (*ta)[i]
	}
}

func countTo1000(n chan int) {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		n <- 1
	}
}

func main() {
	n := make(chan int, 1000);
	a := []int{4,1,-1,4,7,-3,9}
	t := make([]int, len(a))	
	mergeSort(&a, &t, 0, len(a) - 1)
	fmt.Println(a)
	go countTo1000(n)
	sum := 0
	for i := 0; i < 1000; i += <- n {
		sum += 1
	}
	fmt.Println("sum: ", sum)
}