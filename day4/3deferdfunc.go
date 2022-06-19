/*
Go has a special statement called defer that schedules a function call to be run after the function completes.

It keeps our Close call near our Open call so it's easier to understand.
If our function had multiple return statements (perhaps one in an if and one in an else), Close will happen before both of them.
Deferred Functions are run even if a runtime panic occurs.
Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
You can put multiple functions on the "deferred list", like this example.
*/

package main

import (
	"fmt"
)

func first() {
	fmt.Printf("\"first\": %v\n", "first")
}

func second() {
	fmt.Printf("\"second\": %v\n", "second")
}

func main() {
	defer func() {
		first()
	}()
	second()
	for i := 0; i < 10; i++ {
		defer fmt.Printf("i: %v\n", i)
	}
	a := [3]int{1, 2, 3}
	for i, v := range a {
		defer func() {
			a[i] = v
			fmt.Printf("a: %v\n", a)
		}()
	}
	fmt.Printf("a: %v\n", a)
}
