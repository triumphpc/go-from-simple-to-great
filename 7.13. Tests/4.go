/**
Go (Golang) From simple to great. The Complete Developer's Guide.

Can Go automatically convert values to pointers if the method is called on the value itself and not on the value stored in the variable?

func main() {
	value := SuperType(10)
	value.pointerMethod()

	SuperType(15).pointerMethod() // ???

}

func (v *SuperType) pointerMethod() {
	fmt.Println("It's calling here", v)
}

@author Alex Versus 2021
*/
package main

import "fmt"

type SuperType int

func main() {
	value := SuperType(10)
	value.pointerMethod()

	//SuperType(15).pointerMethod() No! It's error!

}

func (v *SuperType) pointerMethod() {
	fmt.Println("It's calling here", v)
}
