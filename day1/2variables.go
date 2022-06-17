//variable - declaration
// A variable is pease of storage containing data temporarily to work with it.
/*
## Naming convention

- A varibale name must begin with letter and can coantain multiple numbers
- A variable can not contain spaces
- A variable starts with lower case letter only accessble inside package and thouse not exported to outside scope of the package
- A variale starts with captital letter exportable outside packge
- Varibales are case sensitive
*/

package main

import "fmt"

// Variable Declaration
var aVariable int

// Variable declaration with initialization
var anotherVariable string = "Hello"

// variable assigned with type inference
var inference = 10

// multi variable declaration with assignment and type inference
var a, x, z = 1, "Heelo", false

// multi block variables declaration
var (
	b string
	c bool
)

//Local variable
func local() {
	var localVariable string
	localVariable = "test"
	fmt.Println(localVariable)
}

/*Local varriable with short hand
short hand will be applicable in local scope
*/
func shorthand() {
	local := "test"
	fmt.Println("shorthand", local)
}

/*
## Zero Values
 If you decalre a variable without assigning any value , Go language assings default/zero value
*/

func zerovalues() {
	var quantity float32
	fmt.Println(quantity)

	var price int16
	fmt.Println(price)

	var product string
	fmt.Println(product)

	var inStock bool
	fmt.Println(inStock)
}

func main() {
	local()
	shorthand()
	zerovalues()
}
