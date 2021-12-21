//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.3: Recipient parameters in methods
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type MyNumber int

type MyPoint *int

// Example for nil
type SomeList struct {
	Key  int
	List *SomeList
}

// Entrypoint
func main() {
	number := MyNumber(2)
	fmt.Println("Original:", number)
	number.Square()
	fmt.Println("After:", number)

	// by point
	number2 := MyNumber(2)
	fmt.Println("Original:", number2)
	number2.SquareByRef() // 4
	fmt.Println("After:", number2)

	// with reference
	pointer := &number2
	pointer.SquareByRef()
	fmt.Println("After with references:", *pointer) // After with references: 8
	// or
	(&number2).SquareByRef() // 16

	// Sum of elements
	el0 := SomeList{2, nil}

	el := SomeList{1, &el0}

	fmt.Println(el.Sum())

}

// Multiplies by two
func (n MyNumber) Square() {
	n *= 2
}

// Multiplies by two by references
func (n *MyNumber) SquareByRef() {
	*n *= 2
}

// Compile error
//func (p MyPoint)  {
//
//}

// Get sum of items in list
func (l *SomeList) Sum() int {
	if l == nil {
		return 0
	}
	return l.Key + l.List.Sum()
}
