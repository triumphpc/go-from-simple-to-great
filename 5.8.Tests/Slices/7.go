/**
Go (Golang) From simple to great. The Complete Developer's Guide.

What output? And how fix it?
-----
func main() {
	st := []string{"", "one", "two", ""}
	fmt.Println(deleteEmpty2(st)) // [one two]
	fmt.Println(st) // ???

}

// Delete empty items from slice
func deleteEmpty2(s []string) []string {
	p := 0
	for _, i  := range s {
		if i != "" {
			s[p] = i
			p++
		}
	}
	return s[:p]
}

-----
@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	st := []string{"", "one", "two", ""}

	//fmt.Println(deleteEmpty2(st)) // [one two]
	//fmt.Println(st) // ???
	st = deleteEmpty2(st)
	fmt.Println(st)

}

// Delete empty items from slice
func deleteEmpty2(s []string) []string {
	p := 0
	for _, i := range s {
		if i != "" {
			s[p] = i
			p++
		}
	}
	return s[:p]
}
