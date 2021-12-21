//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.10: Pointers in methods
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type Animal struct {
	Species string
}

// Entrypoint
func main() {
	var a Animal = Animal{Species: "dog"} // dog

	fmt.Printf("1) %+v\n", a)

	a.setSpecies("cat")
	fmt.Printf("2) %+v\n", a) // cat

}

// Print name of animal
func (a Animal) String() string {
	return a.Species
}

// Set species for Animal
func (a *Animal) setSpecies(s string) {
	a.Species = s
}
