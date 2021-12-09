//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Print the division table from 1 to 10
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var max = 10
	fmt.Printf("%5s", " ")
	for i := 1; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println("")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5s", "-")
	}
	fmt.Println()

	for i := 1; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 1; j <= max; j++ {
			fmt.Printf("%5d", i/j)
		}
		fmt.Println()
	}
}
