/*
An Interface is an abstract type.

Interface describes all the methods of a method set and provides the signatures for each method.
*/

package main

import (
	"fmt"
)

//simple declration of interface

type animal interface {
	eat() bool
	walk() bool
	drink() bool
}

type dog struct {
	age    string
	family string
}

func (d *dog) drink() bool {
	fmt.Println(`Dog can drink`)
	return true
}

func (d *dog) eat() bool {
	fmt.Println(`Dog can eat`)
	return true
}

func (d *dog) walk() bool {
	fmt.Println(`Dog can walk`)
	return true
}

func test1() {
	var a animal
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %T\n", a)
	a = &dog{"asd", "asdasd"}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("age: %v\n", a.(*dog).age)

}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("i is string: %v\n", i)
	case int:
		fmt.Printf("i is int: %v\n", i)
	default:
		fmt.Println(`default`)
	}
}

func main() {
	test1()
	typeSwitch("sdfd")
	typeSwitch(1)
	typeSwitch(1.3)
}
