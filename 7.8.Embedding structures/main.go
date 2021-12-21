//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.8: Embedding structures
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type User struct {
	name string
}

type Customer struct {
	User
	number int
}

// with pointer
type Referal struct {
	*User
}

// Entrypoint
func main() {
	c := Customer{}
	fmt.Printf("s = %#v\n", c) // s = main.Customer{User:main.User{name:""}, number:0}

	c2 := Customer{User{"Alex"}, 666}
	fmt.Printf("s = %#v\n", c2) // s = main.Customer{User:main.User{name:"Alex"}, number:666}
	// the same
	c3 := Customer{User: User{name: "Alex"}, number: 666}
	fmt.Printf("s = %#v\n", c3) // s = main.Customer{User:main.User{name:"Alex"}, number:666}

	// Call getter
	//fmt.Println(c3.Name()) // Alex
	// the same
	fmt.Println(c3.User.Name()) // Alex

	// overloading
	u := User{"User"}
	newC := Customer{User{"Customer"}, 555}

	fmt.Println(u.Name())                   // User
	fmt.Println(newC.User.Name())           // Customer
	fmt.Println(newC.Name(" how are you?")) // Darling Customer how are you?

	newU := &User{"Referal"}
	r1 := Referal{newU}
	r2 := Referal{newU}

	fmt.Println(r1, r2)

}

// Getter for name field
func (u User) Name() string {
	return u.name
}

// Getter for name field
func (c Customer) Name(s string) string {
	return "Darling " + c.name + s
}
