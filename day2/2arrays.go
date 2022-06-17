/**
An array is a data structure that consists of a collection of elements of a single type or simply you can say a special variable, which can hold more than one value at a time. The values an array holds are called its elements or items. An array holds a specific number of elements, and it cannot grow or shrink. Different data types can be handled as elements in arrays such as Int, String, Boolean, and others. The index of the first element of any dimension of an array is 0, the index of the second element of any array dimension is 1, and so on.


**/

package main

import (
	"fmt"
	"os"
	"reflect"
)

func arrays() {
	//Declare arrays
	var arr [5]int
	var stringArr [2]string = [2]string{"a", "b"}
	fmt.Printf("\nValue %v type %s", arr, reflect.ValueOf(arr).Kind())
	fmt.Printf("\nValue %v type %s", stringArr, reflect.TypeOf(stringArr).Kind())

	//Assingment and accessing
	arr[0] = 10
	arr[1] = 12
	fmt.Printf("\nValue %v type %s", arr[1], reflect.ValueOf(arr[1]).Kind())
	ar := [5]float32{1: 2, 3: 5.9}
	fmt.Printf("\nValue %v type %s", ar, reflect.ValueOf(ar[1]).Kind())
	fmt.Println("\nLooping Array")
	for _, v := range ar {
		fmt.Print(v)
	}
	fmt.Println("..")
	for i := 0; i < len(ar); i++ {
		fmt.Printf("%v", ar[i])
	}
	fmt.Println("..")
	for j, _ := range ar {
		fmt.Printf("ar[%v]=%v", j, ar[j])
	}
	fmt.Println("Multidimensional Array")
	m := [4][4]int{{1, 2, 3, 4}, {3, 4, 5}, {2}, {2: 3}}
	fmt.Printf("\n Value %v type %s ", m, reflect.TypeOf(m).Kind())

	fmt.Println("Copy Array")

	aa := [3]int{1, 2, 3}
	bb := aa //data is passed by value
	cc := &aa
	aa[0] = 4
	cc[1] = 5
	(*cc)[2] = 8
	fmt.Println(bb, aa, cc, cc[0], (*cc)[0])
}

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "1":
			arrays()
		}
	}
}
