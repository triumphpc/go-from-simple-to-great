//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a variant function for summing arguments
//
// @author Alex Versus 2021
// www.alexverus.com

package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))    // 6
	fmt.Println(sum(-5, -3, -2)) //-10

}

// Sum all args
func sum(a ...int) (t int) {
	t = 0
	for _, v := range a {
		t += v
	}
	return
}
