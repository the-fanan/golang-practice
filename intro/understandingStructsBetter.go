package main

import (
	"fmt"
)

type General struct {
	Name string
	Num int
}

func (g *General) sayName() {
	fmt.Println(g.Name)
}

func (g *General) sayOtherName() {
	fmt.Println(g.Name, "other name")
}

func (g *General) sayNameDiffSig() {
	fmt.Println(g.Name, "diff signature")
}

func (g *General) updateName(n string) {
	g.Name = n
}

func (g *General) updateNum() {
	g.Num += 1
	fmt.Println(g.Num)
}


type Extender struct {
	General
}

func (e Extender) sayName() {
	fmt.Println("I am from extender", e.Name)
}

func (e *Extender) sayNameDiffSig(s string) {
	fmt.Println(s, "diff signature extender")
}

type Extender2 struct {
	General
}

func (g Extender2) updateName(n string) {
	g.Name = n
}

func (e Extender2) sayName() {
	fmt.Println("I am from extender2", e.Name)
}


type Structs interface {
	sayName()
	updateName(s string)
}

func callUpdateName(s Structs) {
	s.updateName("New Name")
}

func callSayName(s Structs) {
	s.sayName()
}
/**
* When you extend:
* The extender has access to propertis and functionof the embedded struct
* If extender has same function name as embedded struct, it would override
*/
func main() {
	e := Extender{}
	f := &Extender{}
	g := Extender2{}
	e.updateName("Fanan e")
	f.updateName("Fanan f")
	e.sayName()
	e.sayOtherName()
	e.sayNameDiffSig("Dala")//even with different signatures, it sill overrides. Polymorphism is not allowed
	e.updateNum()
	callUpdateName(&e)//e has to be passed by referece because it is not a pointer
	callUpdateName(f)//f is already a pointer to struct Extender
	callUpdateName(g)
	//callSayName(e) -- this would cause and error since it is not passed as a pointer and one or more functions require it be passed as a pointer. Why? Because we cannot 'REREFERENCE' values
	callSayName(&g)
	callSayName(f)
	//this is because the definition for sayName requires a pointer to struct Extender to be passed
	//confusing, that if the function is defined to be passed by value and a reference is passed, it still works. Why? Because in Go, even pointers are passed by value. So, where a value is needed and a pointer is passed, the pointer is DEREFFERENCED -- for structs
	//when it passes by value insead of reference, update functions do not work
}