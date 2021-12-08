//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. What's the output?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	v := 0
	if v := f(); v == 1 {
		fmt.Println("Yep")
	}
	fmt.Println(v)

}

func f() int {
	return 1
}
