/**
## Datatypes
Go is a statically typed language. This means that the type of a variable is known at compile time and type can not be changed.
**/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// int
	var a int = 10
	fmt.Println(reflect.TypeOf(a))
	// float64
	var b float32 = 10.10
	fmt.Println(reflect.TypeOf(b))
	// string
	var c string = "Hello"
	fmt.Println(reflect.TypeOf(c))
	// bool
	var d bool = true
	fmt.Println(reflect.TypeOf(d))
	// complex64
	var e complex64 = 10 + 10i
	fmt.Println(reflect.TypeOf(e))
	// complex128
	var f complex128 = 10 + 10i
	fmt.Println(reflect.TypeOf(f))
	// rune
	var g rune = 'a'
	fmt.Println(reflect.TypeOf(g))
	// byte
	var h byte = 'a'
	fmt.Println(reflect.TypeOf(h))
	// int8
	var i int8 = 10
	fmt.Println(reflect.TypeOf(i))
	// int16
	var j int16 = 10
	fmt.Println(reflect.TypeOf(j))
	// int32
	var k int32 = 10
	fmt.Println(reflect.TypeOf(k))
	// int64
	var l int64 = 10
	fmt.Println(reflect.TypeOf(l))
	// uint
	var m uint = 10
	fmt.Println(reflect.TypeOf(m))
	// uint8
	var n uint8 = 10
	fmt.Println(reflect.TypeOf(n))
	// uint16
	var o uint16 = 10
	fmt.Println(reflect.TypeOf(o))
	// uint32
	var p uint32 = 10
	fmt.Println(reflect.TypeOf(p))
	// uint64
	var q uint64 = 10
	fmt.Println(reflect.TypeOf(q))
	// uintptr
	var r uintptr = 10
	fmt.Println(reflect.TypeOf(r))
	// float32
	var s float32 = 10.10
	fmt.Println(reflect.TypeOf(s))
	// float64
	var t float64 = 10.10
	fmt.Println(reflect.TypeOf(t))
}
