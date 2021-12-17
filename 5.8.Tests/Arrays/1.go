/**
Go (Golang) From simple to great. The Complete Developer's Guide.

What output this code? How change array?

ANSWER:
Arrays in Go are values.
When you pass an array to a function, the array is copied.

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	ar := [3]int{5, 6, 7}
	change(ar)
	fmt.Println(ar) //[5 6 7]

	// convert to slice
	correctChange(ar[:])
	fmt.Println(ar) //[1 2 7]
}

// change some array
func change(a [3]int) {
	a[0] = 1
	a[1] = 2
}

func correctChange(s []int) {
	s[0] = 1
	s[1] = 2
}
