/**
Go (Golang) From simple to great. The Complete Developer's Guide.

The code below scans the directory and initiates panic if it cannot open some directory.
But there is a problem - it catches any panic, even if it was not initiated in the dirScan function.
Rewrite the program so that it displays the panic text and continues to work on scanning directories



func main() {
	// Catch all panic situations
	defer panicCatch()

	// fake panic
	panic("fake panic")

	// scan go directory
	dirScan("/var")


}

// Panic catching
func panicCatch() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)

	if ok {
		fmt.Println(err)
	}
}

// Scan current directory
func dirScan(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			dirScan(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

@author Alex Versus 2021
www.alexversus.com
*/
package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// Catch all panic situations
	defer panicCatch()

	// fake panic
	panic("fake panic")

	// scan go directory
	dirScan("/var")

}

// Panic catching
func panicCatch() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)

	if ok {
		fmt.Println(err)
	} else {
		fmt.Println(p)
		defer panicCatch()
		dirScan("/var")
	}
}

// Scan current directory
func dirScan(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			dirScan(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}
