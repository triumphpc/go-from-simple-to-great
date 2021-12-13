/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Output Duration type in seconds:
---------
ct, _ := time.ParseDuration("1h")

fmt.Printf("%T , %[1]v",ct)
---------

@author Alex Versus 2021
www.alexversus.com
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ct, _ := time.ParseDuration("1h")

	fmt.Printf("There are %.0f seconds in %v.\n", ct.Seconds(), ct)

}
