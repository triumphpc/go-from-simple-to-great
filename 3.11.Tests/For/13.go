//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// List directories up to the current program file
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	st := strings.Split(path, "/")

	for _, p := range st {
		fmt.Println(p, "->")
	}
}
