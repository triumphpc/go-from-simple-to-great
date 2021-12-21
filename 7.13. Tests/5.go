/**
Go (Golang) From simple to great. The Complete Developer's Guide.

Is it possible to write a method for a globally defined type? For example for int:

func (v int) pointerMethod() {
	fmt.Println("It's calling here", v)
}

@author Alex Versus 2021
*/

package main

func main() {

}

// No, it is forbidden
//func (v int) pointerMethod() {
//	fmt.Println("It's calling here", v)
//}
