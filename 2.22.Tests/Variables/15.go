//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Redeclare variable to another value in multiple declaration
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	month, Year := 11, 2021
	month, day := 12, 5

	fmt.Println(day, month, Year)

}
