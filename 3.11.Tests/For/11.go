//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Implement an infinite loop in the program
// 2. At each step of the loop, output an arbitrary character
// 3. After 500 milliseconds, the loop should end on its own
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"time"
)

func main() {
	st := makeTimestamp()
	var p string
	for {
		p += "*"
		fmt.Printf("\r Processing: %s", p)

		if d := makeTimestamp() - st; d > 3 {
			break
		}

	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
