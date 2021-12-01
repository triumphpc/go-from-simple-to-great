// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.7: using Sprintf
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint for program
func main() {
	// We use formatting float and return value
	resultString := fmt.Sprintf("with two digits after the decimal point: %0.2f\n", 1.0/5.0)
	fmt.Println("It's Println", resultString)
}
