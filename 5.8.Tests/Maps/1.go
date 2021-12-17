//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function for comparing two cards.
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	fmt.Println(compare(map[string]int{"A": 1}, map[string]int{"B": 2})) // false
	fmt.Println(compare(map[string]int{"A": 1}, map[string]int{"A": 1})) // true
}

// Compare maps
func compare(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for kx, vx := range x {
		if vy, ok := y[kx]; !ok || vy != vx {
			return false
		}
	}
	return true
}
