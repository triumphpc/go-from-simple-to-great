/**
Go (Golang) From simple to great. The Complete Developer's Guide.

How to change the code to display correct output?
CODE:
func toNegative(myBoolean bool) bool {
	return !myBoolean
}

func main() {
	state := true
	toNegative(state)
	fmt.Println(state) // true

	state = false
	toNegative(state)
	fmt.Println(state) // false
}

OUTPUT:
false
true


@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func toNegative(state *bool) bool {
	return !*state
}

func main() {
	state := true
	toNegative(&state)
	fmt.Println(state) // true

	state = false
	toNegative(&state)
	fmt.Println(state) // false
}
