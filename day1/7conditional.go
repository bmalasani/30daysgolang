package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter two digits")
	scanner.Scan()
	a, ea := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, eb := strconv.Atoi(scanner.Text())

	if ea == nil && eb == nil {

		if a > b {

			fmt.Println("a > b")
		} else {
			fmt.Println("b>a")
		}

	} else {
		panic("error reading numbers")
	}
	ifelsewithdeclaration()
	switchfun()
	switchcoditions()
}

func ifelsewithdeclaration() {
	if day := time.Now(); day.Day() > 15 {
		fmt.Println("'You are in second half'")
	} else if day.Hour() > 12 {
		fmt.Println("AM")
	} else {
		fmt.Println("What")
	}

}

func switchfun() {
	switch time.Now().Day() {
	case 1, 2, 3, 4, 5, 6, 7:
		fmt.Println("First Week")
	case 10, 12, 28, 30:
		fmt.Println("Some Random Day")
	default:
		fmt.Println("Some fucking day")
	}
}

func switchcoditions() {
	day := time.Now().Day()
	switch {
	case day > 15:
		fmt.Println("After 15")
		fallthrough //fallthrough keyword forces excution to next case block
	case day > 17:
		fmt.Println("17th Day")
		fallthrough
	default:
		fmt.Println("fucking day")
	}
}
