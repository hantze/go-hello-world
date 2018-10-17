package main

import "fmt"

type Anu struct {
	Property int
}

func (a *Anu) test1() {
	fmt.Println("anu")
}

type Anu2 struct {
	Anu
}

func (a *Anu2) test1() {
	fmt.Println("anu2")
}

func myTest2(a Anu) {
	fmt.Println("test 2 ", a.Property)
}

func myTest() *Anu2 {
	return &Anu2{
		Anu{
			Property: 12,
		},
	}
}

func main() {
	fmt.Println("Hello World!")
	a := Anu2{}
	a.Anu.test1()       // call from Parent
	a.Anu.Property = 10 // set property
	fmt.Println(a.Anu.Property)
	a.test1() // overriding

	myTest2(a.Anu)

	temp := myTest()               // call myTest
	fmt.Println(temp.Property)     // direct access
	fmt.Println(temp.Anu.Property) // access via parent
}
