//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {

	for {
	run:
		fmt.Println("yep")

		switch {
		case true:
			continue run
		}
	}
}
