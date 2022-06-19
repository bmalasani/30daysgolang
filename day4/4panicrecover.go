/*
The built-in type system of GO Language catches many mistakes at compile time, but unable to check mistakes like an out-of-bounds array, access or nil pointer deference which require checks at run time. GO does not have an exception mechanism you can't throw exceptions. During the execution when Go detects these mistakes, it panics and stops all normal execution, all deferred function calls in that goroutine are executed and finally program crashes with a log message. This log message usually has enough information to analyze the root cause of the problem without running the program repeatedly, so it should always be included in a bug report about a panicking program.
*/

/*Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
 */

package main

import "fmt"

func main() {
	var action int
	defer func() {
		action := recover()
		fmt.Println(action)
	}()
	for {
		fmt.Println("Chooose 3 to panic")
		fmt.Scanln(&action)
		switch action {
		case 3:
			panic("you choosen Panic")
		default:
			fmt.Println("Hello there!")
		}
	}
}
