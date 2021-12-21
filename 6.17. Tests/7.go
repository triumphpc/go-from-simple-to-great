//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a change function that changes the occurrences of the uf func(string) string  received in the first passed argument s string
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(s string) string {
		return "old text " + s
	}

	fmt.Println(change("is's old text this", "this", f)) // is's REPLACED
}

// change input string
func change(current string, s string, uf func(string) string) string {
	return strings.Replace(current, uf(s), "REPLACED", -1)
}
