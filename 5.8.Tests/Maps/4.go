//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Check key "three" exist in map.
// INPUT:
// map[string]int{"one":1, "two":2}
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	m := map[string]uint{"one": 1, "two": 2}

	fmt.Println(m)

	_, ok := m["three"]
	fmt.Println("has key three:", ok)
}
