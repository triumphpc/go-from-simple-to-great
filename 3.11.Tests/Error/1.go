//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Input count of hours from cmd line
// 2. Convert to minutes and output
// 3. Add error handler if input is not number
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
	if len(os.Args) < 2 {
		fmt.Println("Please enter count of hours")
		return
	}

	arg := os.Args[1]

	hours, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		fmt.Printf("error: '%s' is not a integer.\n", arg)
		return
	}

	minutes := hours * 60

	fmt.Printf("%v hours is %v minutes.\n", hours, minutes)
}
