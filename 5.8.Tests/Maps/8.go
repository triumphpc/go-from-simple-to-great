//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write function for merge  maps. The keys and values of second map getting added in first map.
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
)

func main() {
	m1 := map[string]int{"one": 2, "two": 7, "alfa": 1}
	m2 := map[string]int{"one": 5, "five": 5, "alfa": 6}
	merge(m1, m2)

	fmt.Println(m1) // map[alfa:6 five:5 one:5 two:7]
}

func merge(m1, m2 map[string]int) {
	for k, v := range m2 {
		m1[k] = v
	}
}
