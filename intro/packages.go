package main

import (
	"fmt"
	"strings"
)

func capitalizeFirst(s string) string {
	sb := []byte(s)
	return strings.ToUpper(string(sb[0])) + string(sb[1: len(sb)])
}
func main() {
	/**
		* Strings
	*/
	s := "users"
	fmt.Println(strings.Contains("makered", "ake"))
	fmt.Println(strings.Count("makered", "e"))
	fmt.Println(strings.HasPrefix("makered", "mak"))
	fmt.Println(strings.HasSuffix("makered", "red"))
	fmt.Println(strings.Index("makered", "e"))
	fmt.Println(strings.Join([]string{"Make","red"}, "-"))
	fmt.Println(strings.Split("a-b-c-d-e---f-", "-"))
	fmt.Println(strings.ToLower("MAKERS"))
	fmt.Println(strings.ToUpper("users"))
  fmt.Println(capitalizeFirst(s))
	fmt.Println(strings.Repeat("f", 5))
	// func Replace(s, old, new string, n int --- first n srings found) string
	fmt.Println(strings.Replace("fasssxfxxfxffff", "f", "d", 3))
}