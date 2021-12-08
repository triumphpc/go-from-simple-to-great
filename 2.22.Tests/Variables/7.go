//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// How we can print type of variable?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var varInt, varIntSecond = 1, -10
	var varFloat, varFloatSecond = 0.0, 4.15
	var varBool, varBoolSecond = true, false

	fmt.Printf("%v : %T\n", varInt, varInt)
	fmt.Printf("%v : %T\n", varIntSecond, varIntSecond)
	fmt.Printf("%v : %T\n", varFloat, varFloat)
	fmt.Printf("%v : %T\n", varFloatSecond, varFloatSecond)
	fmt.Printf("%v : %T\n", varBool, varBool)
	fmt.Printf("%v : %T\n", varBoolSecond, varBoolSecond)
}
