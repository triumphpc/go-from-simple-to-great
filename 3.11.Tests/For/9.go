//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a cycle that runs twice. Use goto to return to the beginning of the loop
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {

	var j = true
start:
	for i := 0; i < 5; i++ {
		fmt.Println("loop", i)

		if i == 4 && j == true {
			i = 0
			j = false
			goto start
		}
	}
}
