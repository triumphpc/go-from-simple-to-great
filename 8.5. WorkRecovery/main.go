//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 8.5: Emergency situations and work recovery
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint
func main() {

	// recovery example
	//fmt.Println(recover()) // nil

	//if true {
	//	panic("Some panic end")
	//}
	//
	//// recovery not work
	//fmt.Println(recover())

	// example recovery
	first()
	fmt.Println("normal out")

	//defer down()
	//err := fmt.Errorf("that is output error")
	//panic(err)

}

func first() {
	defer second()
	panic("that is panic")
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
	fmt.Println(recover()) //that is panic
}

func down() {
	//p := recover()
	//fmt.Println(p.Error()) // unavailable
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error()) //that is output error
	}
}
