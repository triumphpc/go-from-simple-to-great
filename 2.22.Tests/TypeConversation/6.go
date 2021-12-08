//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Fix error for output : 650
// INPUT:
// var x int8 = 127
// var y int64 = 500
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var x int8 = 127
	var y int64 = 500

	fmt.Println(y + int64(x) + 23)
}
