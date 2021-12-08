//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.18: Command line Args
// go run main.18.go first second
//
// @author Alex Versus 2021
//

package main

import (
	"flag"
	"fmt"
	"os"
)

// Entrypoint
func main() {
	// Flags
	param := flag.String("param", "bar", "some example")
	flag.Parse()
	fmt.Println("param:", *param)

	var args = os.Args
	fmt.Printf("%v : %T\n", args, args)
	fmt.Printf("%v: %T\n", args[0], args[0])
	fmt.Printf("%v: %T\n", args[1], args[1])
	fmt.Printf("%v: %T\n", args[2], args[2])

	var all, sep string
	for i := 1; i < len(os.Args); i++ { // get all slices
		all += sep + os.Args[i] // by index
		sep = ", "
	}
	fmt.Println(all)

}
