//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Develop a program that displays a table with the results of operations.
// 2. The table takes three values as input, for example: go run 12.go <v> <h> <c>,
// where <v> is the number of columns
// <h> is the number of rows
// <c> is a mathematical command,
// possible options are: s - sum, d - division, m - multiplication, sb - subtraction
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
	if len(os.Args) < 4 {
		fmt.Println("Please enter format <v> <h> <c>.")
		return
	}

	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("v not a number.")
		return
	}

	h, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("h not a number.")
		return
	}

	fmt.Printf("%5s", " ")
	for i := 1; i <= h; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println("")
	for i := 0; i <= h; i++ {
		fmt.Printf("%5s", "-")
	}
	fmt.Println()

	for i := 1; i <= v; i++ {
		fmt.Printf("%5d", i)

		for j := 1; j <= h; j++ {
			switch os.Args[3] {
			case "s":
				fmt.Printf("%5d", i+j)
				break
			case "d":
				fmt.Printf("%5d", i/j)
				break
			case "m":
				fmt.Printf("%5d", i*j)
				break
			case "sb":
				fmt.Printf("%5d", i-j)
				break
			default:
				fmt.Println("c - unknown.")
				return
			}
		}
		fmt.Println()
	}
}
