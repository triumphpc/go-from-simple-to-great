//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Declare three constants
// 2. The third constant must be defined by an expression using the first two constants
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

const (
	first, second int8 = 5, 10
	third         int8 = first * second
)

func main() {
	fmt.Println(third)

}
