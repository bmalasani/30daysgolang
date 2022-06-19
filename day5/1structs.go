// - structs can help to create user defined types in go

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

// simple struct

func simple() {
	type Animal struct {
		Category string
	}
	var a Animal
	fmt.Printf("a: %v\n, type: %s \n, kind: %s \n", a, reflect.TypeOf(a).Name(), reflect.TypeOf(a).Kind())
	fmt.Printf("a: %v\n", unsafe.Sizeof(a))

	a = Animal{
		Category: "Mammel",
	}

	fmt.Printf("a: %v\n, type: %s \n, kind: %s \n", a, reflect.TypeOf(a).Name(), reflect.TypeOf(a).Kind())
	fmt.Printf("a: %v\n", unsafe.Sizeof(a))

}

// struct declarations
// - A struct can be defined in global scope / local scope

//Globla struct

type Box struct {
	length, width int
}

// check Global vs local struct
// - Local struct takes precedence
func test1() {
	type Box struct {
		length, width, height int
	}
	var a Box
	a = Box{length: 1}
	fmt.Printf("a: %v\n", a)
}

//declarations
func test2() {
	type Box struct {
		length, width int
	}
	var a Box
	b := Box{length: 1}
	c := &Box{length: 2}
	//with new
	f := new(Box)
	f.length = 12
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("f: %v\n", f)

	//anonymous structs
	g := struct{ len, wid int }{len: 10}
	fmt.Printf("g: %v\n", g)
}

//structs are comparable
func test3() {
	var a = new(Box)
	var b = new(Box)
	fmt.Printf("a==b: %v\n", a == b)
	var aa = Box{length: 10, width: 1}
	var bb = Box{width: 1, length: 10}
	fmt.Printf("aa == bb: %v\n", aa == bb)
	c := &Box{10, 1}
	d := &Box{10, 1}

	//c and d are different pointers
	fmt.Printf("c==d: %v\n", c == d)

	// but c and d has same value
	fmt.Printf("c==d: %v\n", *c == *d)

	//this will be true cause the structs have same values
	fmt.Printf("c==aa: %v\n", *c == aa)

}

//empty struct
func test4() {
	type empty struct{}
	e := empty{}
	fmt.Printf("e: %v\n size: %v\n", e, unsafe.Sizeof(e))
}

//nested structs
func test5() {
	type Parent struct {
		name string
	}
	type Child struct {
		name   string
		parent Parent
	}

	var a = Child{
		name: "child",
		parent: Parent{
			name: "parent",
		},
	}
	fmt.Printf("a: %v\n", a)
}

//promotion

func test6() {
	type Human struct {
		name string
		age  int
	}
	type Cook struct {
		Human  // promoting human as Cook.. now cook can use human props directly
		skills string
	}
	var c = Cook{}
	c.name = "Human promotion"
	c.age = 23
	c.skills = "can cook"
	fmt.Printf("c: %v\n", c)

	type Person struct {
		string     //naming person
		int        // promoting another
		year   int //go cant promote if they have same type
	}
	p := new(Person)
	p.string = "Hello"
	p.int = 23
	p.year = 2022
	fmt.Printf("p: %v\n", p)
}

// struct functions
func test7() {

	//can define func props in structs
	type drive func()
	type car struct {
		modal string
		drive
	}
	c := car{modal: "2022"}
	//like this we cant get c values
	c.drive = func() { fmt.Println("drivving") }
	fmt.Printf("c: %v\n", c)
	c.drive()

}

//struct functions

type A struct {
	name string
}

//struct methods.. these methods cant create in local scopes
func (a A) LongDrive() {
	fmt.Println(a)
}

func test8() {
	a := A{"sdfdf"}
	a.LongDrive()
}

//json
type Employee struct {
	Name string `json: "name"`
}

func test9() {

	empstr := `
	{
		"name":"venakt"
	}`
	var emp = &Employee{}
	json.Unmarshal([]byte(empstr), emp)
	fmt.Printf("emp: %v\n", emp)
	s, _ := json.Marshal(emp)
	fmt.Printf("s: %s\n, %v", s, string(s))
}
func main() {
	simple()
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
}
