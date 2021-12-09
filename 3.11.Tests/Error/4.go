//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Enter year form command line
// 2. Print it's Leap year or not
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter a year number")
		return
	}

	y, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	if isLeapYear(y) {
		fmt.Println(y, "leap year")
	} else {
		fmt.Println(y, "not leap year")
	}
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}
