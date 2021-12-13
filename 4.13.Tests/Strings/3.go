/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Write a program that concatenates all the input arguments of a script into a single string and separates them with commas

Example: go run 3.go go is good
Output: go, is, good

@author Alex Versus 2021
www.alexversus.com
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	lenArg := len(args[1:])
	out := ""

	for i := 1; i <= lenArg; i++ {
		out += args[i] + ", "
	}

	fmt.Println(out[:len(out)-2])
}
