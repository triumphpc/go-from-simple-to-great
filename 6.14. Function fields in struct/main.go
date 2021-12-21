//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.14: Func as field of struct
//
// @author Alex Versus 2021
// www.alexversusu.com

package main

import "fmt"

// Type for full name of user
type FullNameType func(string, string) string

// User struct
type User struct {
	FirstName, LastName string
	FullName            FullNameType
}

// Entrypoint
func main() {
	em := User{
		FirstName: "Alex",
		LastName:  "Versus",
		FullName: func(fn string, ln string) string {
			return fn + " " + ln
		},
	}

	fmt.Println(em.FullName(em.FirstName, em.LastName)) // Alex Versus
}
