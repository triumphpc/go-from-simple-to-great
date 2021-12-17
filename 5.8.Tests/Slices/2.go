//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function code that, without additional memory allocation, will move the array elements forward by 2
// FOR EXAMPLE:
// t2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
// OUTPUT
// t2 := [...]int{2, 1, 9, 8, 7, 6, 5, 4, 3}
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	// revert array
	t2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	rev(t2[:])
	rev(t2[:2])
	rev(t2[2:])
	fmt.Println(t2) // [2 1 9 8 7 6 5 4 3]

}

// Turn items in reverse order
func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
