/**
Go (Golang) From simple to great. The Complete Developer's Guide.

What output of this code? Why?

Go structs are value types. When we assign a struct variable to another struct variable, a new copy of the struct is created.
Likewise, when we pass a struct to another function, the function receives a new copy of the struct.

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	u1 := User{"Name One", 30}
	u2 := u1

	u2.name = "Name Two"
	u2.age = 44

	fmt.Println(u1, u2) // {Name One 30} {Name Two 44}
}
