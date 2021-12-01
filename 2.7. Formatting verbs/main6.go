// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.7: Formatting verbs
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
)

type coordinates struct {
	x, y int
}

// entrypoint
func main() {
	c := coordinates{1, 2}
	// Instance of type coordinates
	fmt.Printf("%v\n", c) // {1 2}
	// Include struct fields name
	fmt.Printf("%+v\n", c) // {x:1 y:2}
	// Go syntax representation of the value. Code snippet reproduced as value
	fmt.Printf("%#v\n", c) // main.coordinates{x:1, y:2}
	// Type of value
	fmt.Printf("%T\n", c) // main.coordinates
	// Print boolean
	fmt.Printf("%t\n", false) // false
	// Print digits
	fmt.Printf("%d\n", 555) // 555
	// Binary representation
	fmt.Printf("%b\n", 123) // 1111011
	// Character corresponding by integer
	fmt.Printf("%c\n", 35) // #
	// Basic formatting for float
	fmt.Printf("%f\n", 99.9) // 99.900000
	// Scientific notation for float
	fmt.Printf("%e\n", 99.9) // 9.990000e+01
	fmt.Printf("%E\n", 99.9) // 9.990000E+01
	// String basic
	fmt.Printf("%s\n", "versus") // "versus"
	// String with double quotes
	fmt.Printf("%q\n", "versus") // ""versus""
	// In base 16 with character per byte of input
	fmt.Printf("%x\n", "versus") // 766572737573
	// Pointer presentation
	fmt.Printf("%p\n", &c) // 0xc000018080
	// Width format with integer
	fmt.Printf("|%5d|%5d|\n", 333, 123) // |  333|  123|
	// Width format with float
	fmt.Printf("|%7.2f|%7.2f|\n", 333.3, 123.123) // | 333.30| 123.12|
	// Width format with float ald left margin
	fmt.Printf("|%-7.2f|%-7.2f|\n", 333.3, 123.123) // |333.30 |123.12 |
	// Return formatted string
	s := fmt.Sprintf("it's %d", 123)
	fmt.Println(s) // it's 123
	// Formatted string to stdout
	fmt.Fprintf(os.Stdout, "%s\n", "www.versus.com")

}
