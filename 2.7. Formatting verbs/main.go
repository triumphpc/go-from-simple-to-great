// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.7: using Println
//
// @author Alex Versus 2021
// www.alexversus.com

package main

// Import packages format
import (
	"fmt"
)

// Entrypoint for program
func main() {
	// Run function Println from fmt package
	fmt.Println(
		"This is Println function exec.",
	)

	const name, age = "Alex", 32
	fmt.Println(name, "is", age, "years old.")

	// It's very loudly
	fmt.Println("It's too many numbers:", 1.0/9.0)

	fmt.Println("About one-third:", 1.0/3.0)

}
