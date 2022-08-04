package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
	"meb.com/greet"
)

func main() {
	s := "Hello"
	greet.Greet(s)
	fmt.Printf("s: %v\n", stringutil.Reverse(s))
}
