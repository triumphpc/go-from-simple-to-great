//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. What's the output?
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	s := "hello go friends!"
	for i := 0; i < len(s); i++ {
		s := s[i]
		if s != '!' {
			s := s
			fmt.Printf("%c", s)
		}
	}
}
