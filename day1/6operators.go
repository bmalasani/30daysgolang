/**
Operators

Arithmatic Operator
+,-,*,/,%,++,--
**/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	//fmt.Println(false+false) //Boolean addition not possible
	fmt.Println(10.5 + 1.6)
	fmt.Println(complex(10, 10) + complex(1, 1))
	fmt.Println(complex(10, 10) + 12)
	fmt.Println("aasds" + "dfgfdg") //string addition
	fmt.Println(10.22 + 1)          //Default conversion?

	var a int64 = 3
	var b = 3.5
	var c = "string"
	var d = "4"

	fmt.Println(b + float64(a))
	fmt.Println(c + strconv.FormatInt(a, 2)) //need to do additional conversion
	num, err := strconv.Atoi(d)              //converting string to int
	if err == nil {
		fmt.Println(num + int(a))
	}

	fmt.Println(b * float64(a))
	fmt.Println(4 / 2)
	fmt.Println(41 % 2)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	//--a
	f := a
	fmt.Println(f)
	f += a
	fmt.Println(f)
	f -= 2
	fmt.Println(f)

	f *= 3
	fmt.Println(f)

	f /= 4
	fmt.Println(f)

	f %= 3

	fmt.Println(f)
	fmt.Println(f == a)
	fmt.Println(a != f)
	fmt.Println(a > f)
	fmt.Println(a < f)
	fmt.Println(a <= f)
	fmt.Println(a >= f)
	fmt.Println(a > f || f > a)
	fmt.Println(a > f && a < f)
	fmt.Println(!(a > f))

	fmt.Println(2 & 3)
	//2= 10 3=11 = 2&3 10
	fmt.Println(2 | 3)
	//2= 10 3=11 = 2|3 11
	fmt.Println(2 ^ 5)
	//2= 010 5=101 = 2^5 111
	fmt.Println(2 << 3)
	//2= 010 2<<3  10000
	fmt.Println(14 >> 2)
	//14= 1110 14>2  0011

}
