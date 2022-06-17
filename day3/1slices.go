package main

import (
	"fmt"
	"reflect"
)

func slices() {
	fmt.Println("\n\nA slice is a flexible and extensible data structure to implement and manage collections of data. Slices are made up of multiple elements, all of the same type. A slice is a segment of dynamic arrays that can grow and shrink as you see fit. Like arrays, slices are index-able and have a length. Slices have a capacity and length property.")

	var aSlice []int              // Empty slice
	var b []int = make([]int, 20) // Declaring slice and assigning space
	fmt.Printf("\nValue %v type %s", aSlice, reflect.TypeOf(aSlice).Kind())
	aSlice = make([]int, 5, 10)
	// we cant push more than length but if we assing another slice with hight size it will be allowed
	for i := 0; i < 20; i++ {
		b[i] = i
	}
	aSlice = b
	fmt.Printf("\nValue %v type %s", aSlice, reflect.TypeOf(aSlice).Kind())

	sliceLiteral := []string{"a", "b", "c", "d"}
	fmt.Printf("\nValue %v type %s", sliceLiteral, reflect.TypeOf(sliceLiteral).Kind())

	var sliceWithNew = new([20]int)[0:5]
	sliceWithNew[0] = 1
	fmt.Printf("\nValue %v type %s", sliceWithNew, reflect.TypeOf(sliceWithNew).Kind())

	//append is used to append items in slice without concern on size

	zeroslice := make([]int, 0)
	zeroslice = append(zeroslice, 0, 1, 2, 3, 4, 5)
	fmt.Printf("\nValue %v type %s", zeroslice, reflect.TypeOf(zeroslice).Kind())

	//access items or subset
	fmt.Printf("\nSingle Value %v sub Array %v", zeroslice, zeroslice[1:])
	zeroslice = append(zeroslice[len(zeroslice)-1:], zeroslice[0:len(zeroslice)-2]...)
	fmt.Printf("\nValue %v type %s", zeroslice, reflect.TypeOf(zeroslice).Kind())

	// built in copy: destination should be bigger than source
	newSlice := make([]int, 10)
	copy(newSlice, zeroslice)
	fmt.Printf("\nValue %v type %s", newSlice, reflect.TypeOf(newSlice).Kind())

}

func main() {
	slices()
}
