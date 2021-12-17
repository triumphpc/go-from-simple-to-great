//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function that will only return non-empty strings in a slice with a string type
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	st := []string{"", "one", "two", ""}
	fmt.Println(deleteEmpty(st)) // [one two]
}

// Delete empty items from slice
func deleteEmpty(s []string) []string {
	p := 0
	for _, i := range s {
		if i != "" {
			s[p] = i
			p++
		}
	}
	return s[:p]
}
