//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Will there be a compilation error in the example below?
// 2. What's the output?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

var pv = "pv"

func main() {
	fmt.Println(fv())
	fv := "fv"
	fmt.Println(fv)
	fmt.Println(pv)
}

func fv() int {
	return 555

}
