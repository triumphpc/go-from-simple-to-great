//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.5: Functions vs methods
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type User struct {
	name string
	age  uint
}

type UserFactory struct{}

// Entrypoint
func main() {
	// Function declaration
	user := NewUser("Alex", 35)
	fmt.Println(*user)

	// call method
	fmt.Println(user.isYoung())

	// Interchangeability
	factory := &UserFactory{}
	u := factory.NewPerson("Nikolas", 15)
	fmt.Println(isYoung(u))

	// Chain of methods
	user2 := &User{}
	user2.withAge(15).withName("Nikola")
	fmt.Println(*user2) // {Nikola 15}

}

// New instance of User
func NewUser(name string, age uint) *User {
	return &User{name, age}
}

// Method for receiver User
func (user *User) isYoung() bool {
	return user.age < 18
}

// `NewPerson` is now a method of a `UserFactory` struct
func (p *UserFactory) NewPerson(name string, age uint) *User {
	return &User{name, age}
}

// `isYoung` is now a function where we're passing the `Person` as an argument
// instead of a receiver
func isYoung(u *User) bool {
	return u.age < 18
}

// Set name for User
func (u *User) withName(name string) *User {
	u.name = name
	return u
}

// Set age for User
func (u *User) withAge(age uint) *User {
	u.age = age
	return u
}
