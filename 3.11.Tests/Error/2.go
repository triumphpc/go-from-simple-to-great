//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Input count of hours from cmd line
// 2. Convert to minutes and output using short statement (shadowing)
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
	if h := os.Args; len(h) < 2 {
		fmt.Println("Please enter count of hours")
		return
	} else if hs, err := strconv.ParseInt(h[1], 10, 64); err != nil {
		fmt.Printf("error: '%s' is not a integer.\n", h)
		return
	} else {
		minutes := hs * 60
		fmt.Printf("%v hours is %v minutes.\n", hs, minutes)
	}
}
