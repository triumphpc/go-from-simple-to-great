//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// What is e value?
// What is iota value for e?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type Even bool

const (
	a = Even(iota%2 == 0)
	b
	c
	d
	e
)

func main() {

	fmt.Println(e) // true, iota = 4

}
