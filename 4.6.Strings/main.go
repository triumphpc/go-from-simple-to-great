//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.6: String
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint
func main() {
	s := "go is amazing"
	fmt.Println(len(s))      // 13
	fmt.Println(s[0], s[12]) // 103 103

	// part of string
	fmt.Println(s[0:2]) // go
	fmt.Println(s[:12]) // go is amazing
	fmt.Println(s[3:])  // is amazing

	// concatenation
	fmt.Println("hello" + s[5:]) // hello amazing go
	s2 := "left"
	t2 := s2
	s2 += " ,right"

	fmt.Println(s2, t2) // left ,right left
	//t2[1] = "up" // error

	str := "string length example"
	fmt.Println(len(str))
	fmt.Println(str[7:])

	// raw string literal
	var rsl string = `First word
Second word`
	fmt.Println(rsl)

	fmt.Println("`foo`")

}
