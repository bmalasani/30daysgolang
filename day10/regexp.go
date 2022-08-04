/*
Package regexp implements regular expression search.

The syntax of the regular expressions accepted is the same general syntax used by Perl, Python, and other languages.
*/

package main

import (
	"fmt"
	"regexp"
)

func regexptest() {
	b := []byte("aaa")
	var eb []byte
	fmt.Println(regexp.Match(`a*`, b))
	fmt.Println(regexp.Match(`.*`, eb))
	fmt.Println(regexp.Match(`a?`, eb))
	fmt.Println(regexp.Match(`a?`, b))
	fmt.Println(regexp.Match(`.`, b))
	fmt.Println(regexp.Match(`.`, eb))
	fmt.Println(regexp.Match(`a+`, eb))
	fmt.Println(regexp.Match(`a+`, b))
	fmt.Println(regexp.Match(`a`, b))
	fmt.Println(regexp.Match(`a`, eb))
}

func main() {
	regexptest()
}
