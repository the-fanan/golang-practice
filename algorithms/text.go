package main

import (
	"fmt"
)

func bruteForce(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	haystackB := []byte(haystack)
	needleB := []byte(needle)
	for i := 0; i < (h - n + 1); i++ {
		j := 0

		for j < n && haystackB[i + j] == needleB[j] {
			j++
		}
		if j == n {
			return i
		}
	}
	return -1
}

/**
 * Boyer-Moore Algorithm
 * Looking-Glass Heuristic: When testing a possible placement of P against T, begin the comparisons from the end of P and move backward to the front of P.
 * Character-Jump Heuristic: During the testing of a possible placement of P within T, a mismatch of text character T[i]=c with the corresponding pattern character P[k] is handled as follows. If c is not contained anywhere in P, then shift P completely past T[i] (for it cannot match any character in P). Otherwise, shift P until an occurrence of character c in P gets aligned with T[i].
 */
func boyerMoore(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	haystackB := []byte(haystack)
	needleB := []byte(needle)
	//for last index dictionary
	l := make(map[string]int)
	for i,v := range needleB {
		l[string(v)] = i
	}

	i := n - 1//traverse haystack
	j := n - 1//traverse needle

	for i < h {
		if haystackB[i] == needleB[j] {
			if j == 0 {
				return i
			}
			i--
			j--
		} else {
			val, ok := l[string(haystackB[i])]
			k := 0
			if ok {
				k = val
			} else {
				k = -1
			}
			if j < k + 1 {
				i += n - j
			} else {
				i += n - (k + 1)
			}
			//reset j
			j = n - 1
		}
	}
	return -1
}
//kmp
func lpsGenerator(needle []byte) []int{
	n := len(needle)

	i := 0
	j := 1
	lps := make([]int, n)
	for j < n {
		if needle[i] == needle[j] {
			lps[j] = i + 1
			i++
			j++
		} else {
			if i == 0 {
				j ++
			} else {
				i = lps[i - 1]
			}
		}
	}

	return lps
}

func main() {
	needle :=  "fanan"
	haystack := "dalafananmoses"
	fmt.Println("Brute: ", bruteForce(haystack,needle))
	fmt.Println("Boyer Moore: ", boyerMoore(haystack,needle))
	arr := make([]int, 10)
	fmt.Println(arr)

	fmt.Println(lpsGenerator([]byte("abcdeabcdebf")))
	//abcdababcdabjkabcd
	//000012123456001234
}