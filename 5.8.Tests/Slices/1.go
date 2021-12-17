//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function that takes a slice as input and reverses the elements in reverse order without using additional memory
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	t1 := []int{1, 3, 5, 7, 0, 9}
	reverse(t1)

	fmt.Println(t1) // [9 0 7 5 3 1]

	// revert array
	t2 := [...]int{9, 5, 3, 1}
	reverse(t2[:])
	fmt.Println(t2)

}

// Turn items in reverse order
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
