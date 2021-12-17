//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 5.5: Struct
//
//@author Alex Versus 2021
// www.alexversion.com

package main

import (
	"fmt"
)

// Entrypoint
func main() {
	// declaration
	var newStruct struct {
		title string
		age   int16
	}

	fmt.Println(newStruct) // nil

	newStruct.title = "This is title"
	fmt.Println(newStruct) // {This is title 0}

	// defined type with base type struct
	type userProfile struct {
		id, phone, balance int
		Name, Address      string
	}

	var newUser userProfile
	fmt.Println(newUser) //{0  0  0}

	// empty struct example
	emptyMap := make(map[string]struct{})
	if _, ok := emptyMap["test"]; !ok {
		emptyMap["test"] = struct{}{} // {} empty body
	}

	// struct literals
	type Point struct {
		X, Y int
	}
	li := 123
	var Yi int
	p := Point{li, Yi}
	fmt.Println(p) //{123 0}

	//init with all fields
	pInit := Point{Y: Yi, X: li}
	fmt.Println(pInit) //{123 0}

	//package one
	//type PackOne struct{x, y int} // not exported
	//
	//package two
	//import one
	//
	//var _ = one.PackOne{x:1, y:2} // error scope

	// fast declaration
	myFastDec := userProfile{
		phone: 111,
		Name:  "test",
	}
	fmt.Println(myFastDec) // {0 111 0 test }

	// field of struct with struct type
	type Address struct {
		Street, City, State, PostalCode string
	}

	var ha Address
	ha.City = "Moscow"
	ha.Street = "Freedom"
	ha.State = "Moscow"
	ha.PostalCode = "123456"
	fmt.Println(ha) // {Freedom Moscow Moscow 123456}

	// defined type with base type struct
	type newUserProfile struct {
		id, phone, balance int
		Name               string
		workAddress        Address
		homeAddress        Address
	}

	user := newUserProfile{
		id:          123,
		Name:        "Ivan",
		homeAddress: ha,
	}
	fmt.Println(user) // {123 0 0 Ivan {   } {Freedom Moscow Moscow 123456}}

	// anonymous fields of struct
	user.Name = "New Ivan"
	user.homeAddress.City = "Saint Petersburg"
	fmt.Println(user)

	type UserProfileWithAnonymous struct {
		id, phone, balance int
		Name               string
		Address
	}

	var userWithAnonymous UserProfileWithAnonymous
	userWithAnonymous.phone = 555
	userWithAnonymous.balance = 10000
	userWithAnonymous.Address.City = "New York"
	userWithAnonymous.Address.Street = "Some street"
	fmt.Println(userWithAnonymous)

	userWithAnonymous.City = "New value"
	userWithAnonymous.Street = "New value for street"
	fmt.Println(userWithAnonymous)

	//type Technology struct {
	//	material string
	//	process string
	//
	//}
	//
	//type Paper struct {
	//	title string
	//	paperType string
	//	technology Technology
	//}
	//
	//type Book struct {
	//	paper Paper
	//	name string
	//}

	type Technology struct {
		material string
		process  string
	}

	type Paper struct {
		title     string
		paperType string
		Technology
	}

	type Book struct {
		Paper
		name string
	}

	var book Book
	book.name = "My book"
	book.title = "Book title"
	book.paperType = "Paper type"
	book.material = "Semi-cellulose"
	book.process = "Sorting and packing"
	fmt.Println(book) // {{Book title Paper type {Semi-cellulose Sorting and packing}} My book}

	// error compilation
	//book = Book{"My book", "Some field value"...} // error
	//book = Book{name: "Book title", paperType : "New value", process: "Name of process"} // error

	// correct forms
	book = Book{
		Paper{
			"Paper Title",
			"Paper type",
			Technology{
				"Material",
				"Process"},
		},
		"Name",
	}

	fmt.Println(book)

	book = Book{
		name: "Book name",
		Paper: Paper{
			title: "Title of paper",
			Technology: Technology{
				material: "Material from technology",
			},
		},
	}
	fmt.Println(book)
	fmt.Printf("%v\n", book)
	fmt.Printf("%#v\n", book)

	// compare struct
	type Compare struct {
		X, Y int
	}
	x := Compare{5, 5}
	y := Compare{6, 5}
	z := Compare{5, 5}
	fmt.Println(x.X == y.X && x.Y == x.Y)
	fmt.Println(x == y) // false
	fmt.Println(x == z) // true

	// use in map as key
	type ForKey struct {
		int // anonymous field
		string
	}

	m := make(map[ForKey]int)
	fmt.Println(m)
	m[ForKey{123, "st"}] = 555
	fmt.Printf("%#v", m) // map[main.ForKey]int{main.ForKey{int:123, string:"st"}:555}

}
