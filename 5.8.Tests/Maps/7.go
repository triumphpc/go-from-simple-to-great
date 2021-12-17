//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write function sort map values for map[string]int
//
// @author Alex Versus 2021
//

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"one": 2, "two": 7, "alfa": 1}
	m = sortByVal(m)

	fmt.Println(m) // map[alfa:1 one:2 two:7]
}

func sortByVal(m map[string]int) (sm map[string]int) {
	values := make([]int, 0, len(m))
	sm = make(map[string]int)

	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)

	for _, v := range values {
		for mk, mv := range m {
			if mv == v {
				sm[mk] = v
				break
			}
		}
	}
	return sm
}
