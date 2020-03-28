package main

import (
	"fmt"
)

func main() {
	//Arrays
	var x [5] float64
	x[0] = 3
	x[1] = 10
	x[2] = 5
	x[3] = 8
	x[4] = 23

	var total float64 = 0
	for _,value := range x {
		total += value
	}
	fmt.Println(total/float64(len(x)))

	//Slices
	//var c [] float64 //slice with length 0
	//d := make([] float64, 5)//create a slice with an associated array of length 5
	//e := make([] float64, 5, 10)//create a slice with length 5, associatedwith an array of length 10
	// slcies are extendable but not beyond the length of the array they are associated with
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	slice3 := make([]int, 2)
	copy(slice3, slice1)
	//copy(TARGET, SOURCE)
	fmt.Println(slice1, slice3)

	// Maps
	m := make(map[string]int)
	m["fan"] = 3
	m["an"] = 2
	fmt.Println(m)
	delete(m, "an")
	fmt.Println(m)

	value, status := m["fan"]
	fmt.Println(value, status)

	//Exercise: min number
	l := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
		}
	var s int
	for _, value := range l {
		if s == 0 {
			s = value
		} else if s > value{
			s = value
		}
	}
	fmt.Println("Min number in list is: ", s)
}