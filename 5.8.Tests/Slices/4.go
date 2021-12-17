/**
Go (Golang) From simple to great. The Complete Developer's Guide.

Compare two slices with int type

FOR EXAMPLE:
s1 := []int{1,2,3}
s2 := []int{2,1,3}

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{2, 1, 3}
	s3 := []int{1, 2, 3}

	fmt.Println(equal(s1, s2))
	fmt.Println(equal(s1, s3))
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true

}
