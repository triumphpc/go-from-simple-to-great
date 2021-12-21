//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// What output of this code?
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	a(5, 5)        // 10
	b(3, 3)        // 9
	c(false)       // true
	d("Repeat", 3) // Repeat\nRepeat\nRepeat\n
	a(7, 3)        // 10
	b(1, 2)        // 2
	d("no", 2)     // no\nno
}

func a(a int, b int) {
	fmt.Println(a + b)
}

func b(a int, b int) {
	fmt.Println(a * b)
}

func c(a bool) {
	fmt.Println(!a)
}
func d(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Println(a)
	}
}
