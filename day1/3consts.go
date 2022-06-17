/**
 A constant represents fixed value in memory.

 Constants can be declare in local and globale scope
 constnst cant be used with shot variable declartion
**/

package main

import "fmt"

const CONST string = "a"
const CONST3 = "12"
const (
	CONST1 string = "b"
	CONST2 int    = 12
)

func main() {

	const CONST4 = false
	fmt.Println(CONST, CONST1, CONST2, CONST3, CONST4)
}
