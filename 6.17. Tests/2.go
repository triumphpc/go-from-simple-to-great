//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// What output of this code?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	fmt.Println(printThis(10, 5))
}

func printThis(w float64, h float64) float64 {
	area := w * h
	return area / 10.0
	//fmt.Println(area / 10.0) // errer missing return at end of function
}
