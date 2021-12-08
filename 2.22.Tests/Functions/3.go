//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Assign variables from func returns
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	i, s := returnSmth()
	fmt.Print(i, s)
}

/*
Func return different types
*/
func returnSmth() (int, string) {
	return 123, "yep"

}
