//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Update code from practice 4.go. Add days count arg
// 2. Print the months of the specified year in which the specified number of days
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
	if len(os.Args) != 3 {
		fmt.Println("Enter a year number and days count")
		return
	}

	y, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	d, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%q is not a valid days.\n", os.Args[2])
		return
	}

	if isLeapYear2(y) && d == 29 {
		fmt.Println("february")
	} else {
		if d == 30 {
			fmt.Println("april", "june", "september", "november")
		} else if d == 31 {
			fmt.Println("january", "march", "may", "july", "august", "october", "december")
		} else if d == 28 {
			fmt.Println("february")
		} else {
			fmt.Println("Unknown months")
		}
	}
}

func isLeapYear2(year int) bool {
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
