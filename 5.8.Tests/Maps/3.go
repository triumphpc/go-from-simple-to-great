//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Delete third item of map
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	m := make(map[string]uint)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m) // map[one:1 three:3 two:2]

	delete(m, "two")

	fmt.Println(m) // map[one:1 three:3]
}
