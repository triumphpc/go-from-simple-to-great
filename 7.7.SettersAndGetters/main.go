//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.7: Setters and Getters
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"github.com/triumphpc/go-from-simple-to-great/go-from-simple-to-great/7.7.SettersAndGetters/user"
	"log"
)

// Entrypoint
func main() {
	u := user.User{}
	err := u.SetName("Alex")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	err = u.SetSecondName("Versus")
	if err != nil {
		log.Fatal(err)
	}
	err = u.SetAge(18)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)

	//u.Age = 30
	//fmt.Println(u)

	fmt.Println(u.Name())
	fmt.Println(u.Age())

}
