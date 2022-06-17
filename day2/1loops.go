package main

import (
	"fmt"
)

func forlooop() {
	for k := 0; k < 10; k++ {
		fmt.Println(k)
	}
	for i := 0; ; i++ {
		fmt.Println(i)
		if i < 10 {
			break
		}
	}
	j := 1
	for j < 20 {
		fmt.Println(j)
		j++
	}
	for {
		fmt.Println(j)
		j++
		if j < 100 {
			break
		}
	}
}
func forlooopArray() {
	var arr [5]int = [5]int{1: 10}
	for a, v := range arr {
		fmt.Println(a, v)
	}
}

func forlooopMap() {
	var arr map[string]int = make(map[string]int)
	arr["a"] = 1
	arr["b"] = 2
	arr["c"] = 3
	for a, v := range arr {
		fmt.Println(a, v)
	}
}

func main() {
	forlooop()
	forlooopArray()
	forlooopMap()
}
