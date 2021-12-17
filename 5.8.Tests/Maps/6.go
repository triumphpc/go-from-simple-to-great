//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write function sort map keys for map[string]int
//
// @author Alex Versus 2021
//

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"one": 1, "two": 2, "alfa": 3}
	m = sortByKey(m)

	fmt.Println(m) // map[alfa:3 one:1 two:2]
}

func sortByKey(m map[string]int) (sm map[string]int) {
	keys := make([]string, 0, len(m))
	sm = make(map[string]int)

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		sm[k] = m[k]
	}
	return sm
}
