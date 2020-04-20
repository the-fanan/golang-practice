package main

import (
	"fmt"
)

type Tree struct {
	Left *Tree
	Right * Tree
	Value int
}

func (t *Tree) traverseDepthFirst() {
	fmt.Println(t.Value)
	fmt.Println("/", "\\")
	if t.Left != nil {
		fmt.Print(t.Left.Value," ")
	}

	if t.Right != nil {
		fmt.Print(t.Right.Value," \n")
	}
	if t.Left != nil {
		t.Left.traverseDepthFirst()
	}

	if t.Right != nil {
		t.Right.traverseDepthFirst()
	}
}

func (t *Tree) insert(n int) {
	if n > t.Value {
		i := t.Value 
		t.Value = n
		if t.Left == nil && t.Right == nil { 
			t.Left = &Tree{Value: i}
		} else if t.Left != nil && t.Right == nil && t.Left.Value < i {
			t.Right = &Tree{Value: i}
		} else if t.Left != nil && t.Right == nil && t.Left.Value > i{
			t.Right = &Tree{Value: t.Left.Value}
			t.Left.Value = i
		} else {
			//both left and right are not nil
			if t.Left.Value < i && i < t.Right.Value {
				t.Left.insert(i)
			} else {
				t.Right.insert(i)
			}
		}
	} else {
		if t.Left == nil && t.Right == nil {
			t.Left = &Tree{Value: n}
		} else if t.Left != nil && t.Right == nil && t.Left.Value < n {
			t.Right = &Tree{Value: n}
		} else if t.Left != nil && t.Right == nil && t.Left.Value > n{
			t.Right = &Tree{Value: t.Left.Value}
			t.Left.Value = n
		} else {
			//both left and right are not nil
			if t.Left.Value < n && n < t.Right.Value {
				t.Left.insert(n)
			} else {
				t.Right.insert(n)
			}
		}
	}
}

func main() {
	t := Tree{Value: 8}
	t.insert(10)
	t.insert(13)
	t.insert(11)
	t.insert(20)
	t.insert(9)
	t.insert(6)
	t.insert(3)
	t.insert(30)
	t.insert(100)
	t.insert(12)
	t.traverseDepthFirst()
}