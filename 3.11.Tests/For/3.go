//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. We have loop:
// for i := 1; i < 10; i = i++ {}
// 2. If current sum in iteration > 10 - add break from for loop
// 3. Sum only even number
// 4. Output sum value of i in every iteration
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var s = 0
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		s += i

		if s > 10 {
			break
		}
	}

	fmt.Println("Sum:", s)
}
