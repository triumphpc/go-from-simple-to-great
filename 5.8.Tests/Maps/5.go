//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// What output? Why?
//
// Answer: m and m2 equal. When you assign a map to a new variable, they both refer to the same underlying data structure.
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	m := map[string]uint{"one": 1, "two": 2}

	m2 := m

	m2["three"] = 3

	fmt.Println(m, m2) // map[one:1 three:3 two:2] map[one:1 three:3 two:2]
}
