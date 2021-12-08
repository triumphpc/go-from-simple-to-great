//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write Celsius to Fahrenheit Temperature Converter
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	// in fahrenheit
	var f = 212.

	// in celsius
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point of water: %g°F or %g°C\n", f, c)
}
