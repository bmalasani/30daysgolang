/*

A function is a group of statements that exist within a program for the purpose of performing a specific task. At a high level, a function takes an input and returns an output.

Function allows you to extract commonly used block of code into a single component.

The single most popular Go function is main(), which is used in every independent Go program.
*/
/*
A name must begin with a letter, and can have any number of additional letters and numbers.
A function name cannot start with a number.
A function name cannot contain spaces.
If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
function names are case-sensitive (car, Car and CAR are three different variables).
*/

package main

import "fmt"

//simple function

func simple() {
	fmt.Println("It is a simple function")
}

//simple function with parameters and return value
//Note: we cant use polymorpism in go as like other languages
func simplearg(arg string) string {
	fmt.Println("Simple func with arg" + arg)
	return arg
}

//go functions are case sensitive means function Simple, simple , SIMPLE are different

// A function/Variable with Capitallized string is public and default exported from the package
func Simple() string {
	return "Simple"
}

// return can be named in go lang

func namedReturn() (PI float32) {
	PI = 22 / 7
	return
}

// go function can return multiple returns

func multiReturns() (int, int) {
	a := 1
	b := 2
	return a, b
}

func namedmultiReturns() (a int, b int) {
	a = 1
	b = 2
	return
}

//pass by value and pass by reference

func passByValue(a int) {
	a = 10
	fmt.Println(a)
}

func passByRefe(a *int) {
	*a = 10
	fmt.Println(*a)
}

//anonymous function is function without name and can asign to a variable
func anonymousInside() func() int {
	return func() int {
		return area(12, 2) //calling anonymous function
	}
}

var area = func(a int, b int) int {
	return a * b
}

//user defined func type

type RETURNINT func(int) int

//variadic functions

func main() {
	simple()
	simplearg("Hello")
	fmt.Println(Simple())
	PI := namedReturn()
	fmt.Println(PI)
	a, b := multiReturns()
	fmt.Println(a, b)
	c, d := namedmultiReturns()
	fmt.Println(c, d)
	f := 1
	passByValue(f)
	fmt.Println(f)
	passByRefe(&f)
	fmt.Println(f)
	fmt.Println(anonymousInside()())
	func(s int) {
		fmt.Println(s)
	}(10)
	var afunc RETURNINT
	Local := 11
	afunc = func(i int) int {
		return i * Local
	}
	fmt.Println("Clousure example", afunc(12))
}
