// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.10: Logical operators
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint for program
func main() {
	// Logical operators
	fmt.Println((18 > 17) && ("start of planting" == "start of planting") && (0 < 3)) // three operators in true
	// it's true
	fmt.Println((5 == 5) || (17 < 5))
	//it's false
	fmt.Println(!(1 < 2))

	// One original expression is false
	fmt.Println((0.7 > 0.4) && (1.3 < -5.3))
	// Expressions is false
	fmt.Println(!((12.2 == 2.2) || (2.4 != 2.4)))
	// It's false
	fmt.Println(!(-5.4 <= 0.0))
	// it's true
	fmt.Println((true || (5 > 4) || (-4 > 4)))
	// It's true
	fmt.Println(!!(false || (5 > 4) && (-4 < 4)))

}
