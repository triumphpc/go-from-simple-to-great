//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// What output?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	a1 := [5]string{"One", "Two", "Three"}
	a2 := a1

	a2[1] = "New"

	fmt.Println(a1) // [One Two Three  ]
	fmt.Println(a2) // [One New Three  ]
}
