//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function that will return the minimum value from the array values. As an argument, the function takes an array [...] uint
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	ar := [...]uint{5, 6, 10, 56}
	fmt.Println(min(ar[:])) // 5

}

// Return min value from array
func min(ar []uint) (mv uint) {
	for _, v := range ar {
		if mv == 0 || v < mv {
			mv = v
		}
	}
	return mv
}
