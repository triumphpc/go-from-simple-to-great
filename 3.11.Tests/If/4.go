//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Simplify the code below
// 2. Code must output number of args
//
// ------------------------
// args := os.Args
//
// if len(args) < 2 {
// 	fmt.Println("Please enter args")
// } else if len(args) < 3 {
// 	fmt.Printf("There is only one args param: %q\n", args[1])
// } else if len(args) == 3 {
// 	fmt.Printf("You have two args: \"%s %s\"\n", args[1], args[2])
// } else {
// 	fmt.Printf("There are %d args\n", len(args))
// }
// ------------------------
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l > 0 {
		if l == 1 {
			fmt.Printf("There is only one args param: %q\n", args[1])
		} else if l == 2 {
			fmt.Printf("You have two args: \"%s %s\"\n", args[1], args[2])
		} else {
			fmt.Printf("There are %d args\n", len(args))
		}
	} else {
		fmt.Println("Please enter args")
	}
}
