package main

import (
	"fmt"
)

type Size int
type comparable interface {
	~int | ~string
}

func print[T comparable](v T) {
	fmt.Println(v)
}

func write[T comparable](v chan T, a T) {
	v <- a
}

func main() {
	print(9)
	var s Size = 90
	print(s)
	print("hgjggs")
	ch := make(chan int, 2)
	go write(ch, 3)
	ch <- 10
	fmt.Printf("ch: %v\n", <-ch)
	fmt.Printf("ch: %v\n", <-ch)

}
