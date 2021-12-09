//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write a function for raising a number to a power and compare with third number
// 2. Example run program: go run 6.go 2 3 5
// 3. Output:
// 8
// 8 > 5
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if n := os.Args; len(n) < 4 {
		fmt.Println("Please enter in format 6.go number power compare_number ")
	} else if num, err := strconv.ParseFloat(n[1], 64); err != nil {
		fmt.Println("Unknown num", num)
	} else if pow, err := strconv.ParseFloat(n[2], 64); err != nil {
		fmt.Println("Unknown power", pow)
	} else if com, err := strconv.ParseFloat(n[3], 64); err != nil {
		fmt.Println("Unknown compare", com)
	} else {
		v := math.Pow(num, pow)
		fmt.Println(v)
		if v > com {
			fmt.Println(v, ">", com)
		} else {
			fmt.Println(v, "<", com)
		}
	}
}
