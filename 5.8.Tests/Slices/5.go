/**
Go (Golang) From simple to great. The Complete Developer's Guide.

What is output of this code?
-----
s1 := []string{"st1", "st2"}
	s2 := append(s1, "s3", "s4")
	s3 := append(s2, "s5", "s6")
	s4 := append(s3, "s7", "s8")

	s2[0] = "newst1"
	s3[0] = "newst2"
-----

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	s1 := []string{"st1", "st2"}
	s2 := append(s1, "s3", "s4")
	s3 := append(s2, "s5", "s6")
	s4 := append(s3, "s7", "s8")

	s2[0] = "newst1"
	s3[0] = "newst2"

	fmt.Println(s1, s2, s3, s4) // [st1 st2] [newst1 st2 s3 s4] [newst2 st2 s3 s4 s5 s6] [newst2 st2 s3 s4 s5 s6 s7 s8]
}
