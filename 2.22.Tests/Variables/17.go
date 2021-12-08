//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Swap two variables
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	f, s := "first", "second"
	f, s = s, f
	fmt.Println(f, s)
}
