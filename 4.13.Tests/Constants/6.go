/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Fixe this code for correct answer:
-------------
const courseDuration, additionalMinutes = 5, 5

result, _ := time.ParseDuration("10h")
result += additionalMinutes

fmt.Printf("%s later...\n", result*courseDuration)

-------------
OUTPUT:50h5m0s

@author Alex Versus 2021
www.alexversus.com
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	const courseDuration, additionalMinutes = 5, 5

	result, _ := time.ParseDuration("10h")
	result += time.Minute + additionalMinutes

	fmt.Printf("%s later...\n", result*courseDuration)
}
