//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function to sum all the items of an array
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	ar := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(ar)) // 15
}

// Sum of items for arrays
func sum(a [5]int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}
