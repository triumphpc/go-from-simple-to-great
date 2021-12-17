//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write compare func for struct
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type User2 struct {
	name string
	age  int
}

func main() {
	u1 := User2{"Name One", 30}
	u2 := u1
	u3 := User2{"Name Two", 30}

	fmt.Println(compare(u1, u2)) // true
	fmt.Println(compare(u1, u3)) // false
}

// Compare structs
func compare(s1, s2 User2) bool {
	return s1 == s2
}
