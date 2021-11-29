// Go. From simple to great. Junior level.
// Lesson 2.4. Naming convention
//
// Alex Versus (www.alexversus.com). 2022.
//
package main

import (
	"fmt"
	"github.com/mailru/easyjson/parser"
	"io"
)

func main() {
	// normal name for variable
	var nameLocal string
	// normal name for global variable of my package
	var FullName string
	// case matter
	var firstLetter string
	// it's different variable
	var FirstLetter string

	// wrong name
	//var 2SomeError float32
	// wrong call function
	//fmt.someFunc()

	// recommendations for naming
	var CamelCase string
	// abbreviation
	var maxSizeOfWorld int8

	fmt.Println(nameLocal, FullName, CamelCase, maxSizeOfWorld, firstLetter, FirstLetter)
}

// normal name for function
func runAction() {
	// the meaning is obvious
	var ml int // ml -> max length

	// wrong, to verbose
	var memorySpent int
	// good, size of memory
	var s int
	// good, in large scope
	var orderClosed bool

	fmt.Println(ml, memorySpent, s, orderClosed)

}

// normal name for Global function of my package
func GiveSomeInfo() {
	// Getters and setters
	price := obj.Price() // getter
	if price != oldPrice {
		obj.SetPrice(price) // setter
	}
}

// NOT IDIOMATIC Unnecessarily verbose
func ParseFile(file *File, currentFormat []format) (data string, err error) {
	if file.empty() {
		file.Close()
	}

	data = parser.Parse(
		format,
		file)
	return data, nil
}

// IDIOMATIC Concise and idiomatic
func NewParseFile(f *File, fmt []format) (d string, err error) {
	if f.empty() {
		f.Close()
	}

	d = parser.Parse(
		fmt,
		f)
	return d, nil
}

// Abbreviation
var s string      // string
var i int         // index
var num int       // number
var msg string    // message
var v string      // value
var val string    // value
var fv string     // flag value
var err error     // error value
var args []string // arguments
var seen bool     //has seen?
var parsed bool   // parsing ok?
var buf []byte    //buffer
var off int       //offset
var op int        //operation
var opRead int    // read operation
var l int         //length
var n int         //number or number of
var m int         //another number
var c int         //capacity
var a int         //array
var r rune        //rune
var sep string    //separator
var src int       // source
var dot int       // destination
var b byte        //byte
var b []byte      //buffer
var buf []byte    //buffer
var w io.Writer   //writer
var r io.Reader   //reader
var pos int       //position
