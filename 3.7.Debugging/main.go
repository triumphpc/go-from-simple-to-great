// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 3.7: Debugging
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

// Entrypoint
func main() {
	// for loop
	var i int = 0
	for {
		i++

		// remainder of the division
		if i%2 == 0 {
			continue
		}

		if i < 10 {
			println("loop", i)
		}

		// break from loop
		if i == 7 {
			break
		}
	}
}
