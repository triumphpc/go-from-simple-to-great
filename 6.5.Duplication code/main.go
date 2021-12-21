//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.5: Duplication
//
//@author Alex Versus 2021
// www.alexversion.com

package main

import "fmt"

// Entrypoint
func main() {
	var w, h, a float64
	w = 4.2
	h = 3.0
	a = w * h

	fmt.Println(a/10.0, "pallets needed")

	// call function
	pallets(w, h)
}

func pallets(w, h float64) float64 {
	a := w * h
	fmt.Println(a/10.0, "pallets needed")
	return a
}
