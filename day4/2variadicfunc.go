/*A variadic function is a function that accepts a variable number of arguments. In Golang, it is possible to pass a varying number of arguments of the same type as referenced in the function signature. To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which shows that the function may be called with any number of arguments of this type. This type of function is useful when you don't know the number of arguments you are passing to the function, the best example is built-in Println function of the fmt package which is a variadic function.

 */

package main

import (
	"fmt"
	"strings"
)

//simple variadic func

func sum(a ...int) (sum int) {
	for _, v := range a {
		sum = sum + v
	}
	return
}

// we can use normal parameters in variadic function but variadic params should be last.. also we cant have two variadic variabls

func testv(b float32, d int, s ...string) {
	fmt.Printf("strings.Join(s, \"\"): %v\n", strings.Join(s, ""))
	fmt.Printf("d: %v\n", d)
	fmt.Printf("b: %v\n", b)
}

//pass different types of params in variadic

func differentVariadic(i ...interface{}) {
	for _, v := range i {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	fmt.Printf("sum of 12,3,4 = %v", sum(12, 3, 4))
	fmt.Printf("sum of 12,1,3,5,3,4 = %v", sum(12, 1, 3, 5, 3, 4))
	testv(2.2, 2, "Sdf", "sdfsdf")
	differentVariadic(2.2, 2, "Sdf", "sdfsdf")
}
