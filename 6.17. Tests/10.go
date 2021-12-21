//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function that will change the structure field. Pass the structure into a function by reference
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type myType struct {
	field string
}

func main() {
	v := myType{"text here"}
	fmt.Print(v)

	changeType(&v)
	fmt.Print(v) //}{new value}

}

// Change field func
func changeType(t *myType) {
	t.field = "new value"
}
