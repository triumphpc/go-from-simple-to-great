//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write verbose output for subtracting numbers entered on the command line.
// Example: go run 10.go 100 3 5 10 -6
// 100 - 3 - 5 - 10 - -6 = 88

// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var r = 0

	for i, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Not a number.", err)
			return
		}
		if i > 0 {
			r -= val
			fmt.Print("- ", val, " ")
		} else {
			r += val
			fmt.Print(val, " ")
		}

	}
	fmt.Print("= ", r, "\n")
}
