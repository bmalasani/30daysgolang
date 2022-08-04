/*

Concurrency in Golang is the ability for functions to run independent of each other. Goroutines are functions that are run concurrently. Golang provides Goroutines as a way to handle operations concurrently.

New goroutines are created by the go statement.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		go fmt.Printf("i: %v\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	wg.Add(1)
	fmt.Println("Start")
	go test()
	fmt.Println("End")
	for i := 0; i < 10; i++ {
		go fmt.Printf("i: %v\n", i)
	}
	wg.Wait()
	fmt.Println("After WG")
}
