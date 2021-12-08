//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Declare a variable in the package-scope
// 2. Print it's variable in function scope
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// package scope
var i, b, _ = 1, true, false

func main() {
	fmt.Println(i, b)
}
