//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Get from arg string value and convert it to int64.
// Example: run go 11.go 9999999999
// Output: int64 (9999999999)
//
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
	v, _ := strconv.ParseInt(os.Args[1], 10, 64)

	fmt.Printf("%T (%[1]v)\n", v)
}
