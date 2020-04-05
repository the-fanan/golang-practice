package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
		* Arrays
	*/
	a := [5]int{1,2,3,4,5}
	fmt.Println(a, len(a))//len() gives length of array
	//create a multi-dimensional array
	m := [2][2]int{{1,2}, {2,4}}
	v := m[1]
	fmt.Println(m)
	fmt.Println(m[0])
	fmt.Println(m[0][1])
	fmt.Println(v)
	fmt.Println(v[1])
	/**   --- SHORTCOMINGS OF ARRAY --
		* Go arrays cannot be expanded
		* They are passed by value rather than by reference hence, they cannot be mutated within functions and use up a lot of memory
	*/

	/**
		* Slices
	*/
	s := make([]int, 10)//create a slice with 10 slots. Each slot initialized with 0
	fmt.Println(s)
	fmt.Println("capacity:", cap(s), "length:", len(s))
	s = append(s, -5)//append a value to s
	fmt.Println("capacity:", cap(s), "length:", len(s))//capacity doubled when length becomes greater than it
	fmt.Println(s)
	s2 := s[5:11]//create new slice from part of old slice
	fmt.Println(s2)
	s3 := make([]int, 4, 8)//create a slice with length 4 and capacity 8
	fmt.Println("capacity:", cap(s3), "length:", len(s3))
	//copying slices
	a1 := []int {2, 3, 4, 8, 10} //create a slice and populate
	a2 := []int {-2, -5, -6}
	copy(a1,a2)
	fmt.Println(a1)//notice, since slices are of unequal length only first 3 elements in a1 are changed
	//arrays cannot be copied into slices.
	//slices ca be multidimentional
	ms := make([][]int, 4)
	fmt.Println(ms)
	ms[0] = append(ms[0],4)
	fmt.Println(ms)
	//sorting
	sS := []int{2,3,1,8,5}
	fmt.Println(sS)
	sort.Slice(sS, func (i,j int) bool{
		return sS[i] < sS[j]
	})
	fmt.Println(sS)
	sort.Slice(sS, func (i,j int) bool{
		return sS[i] > sS[j]
	})
	fmt.Println(sS)

	/**
		* Maps
	*/
	M := make(map[string]int)
	M["biology"] = 80
	M["chemistry"] = 85
	M["physics"] = 90
	M["mathematics"] = 100
	M["literature"] = 68
	fmt.Println(M)
	delete(M, "literature")
	fmt.Println(M)
	//traverseing map
	for k,v := range M {
		fmt.Printf("key: %s, value: %d%%\n",k,v)
	}
}