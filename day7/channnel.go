/*
Go provides a mechanism called a channel that is used to share data between goroutines. When you execute a concurrent activity as a goroutine a resource or data needs to be shared between goroutines, channels act as a conduit(pipe) between the goroutines and provide a mechanism that guarantees a synchronous exchange.
*/

//A channel is created by the make function, which specifies the chan keyword and a channel's element type.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sum(ch, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	time.Sleep(time.Second * 1)
}

func sum(ch chan int, len int) {
	sum := 0
	for i := 0; i < len; i++ {
		sum += <-ch
	}
	fmt.Printf("Sum: %d\n", sum)
}
