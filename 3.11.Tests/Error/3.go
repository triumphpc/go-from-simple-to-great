//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Enter a number from 0 to 100 from the command line
// 2. If it is greater than 80 - the script displays "Good result"
// 3. If from 40 to 80 - the script displays "Normal"
// 4. If less than 40 - crypt output "Bad result"
// 5. If entered 100 - displays "Super"
// 6. Add a check on the erroneous input
//
// @author Alex Versus 2021
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Requires number")
		return
	}

	num, err := strconv.Atoi(os.Args[1])

	if err != nil || num < 0 {
		fmt.Printf(`Wrong number: %q`+"\n", os.Args[1])
		return
	} else if num > 80 {
		if num >= 100 {
			fmt.Println("Super")
		} else {
			fmt.Println("Good result")
		}
	} else if num <= 80 && num >= 40 {
		fmt.Println("Normal")
	} else if num < 40 {
		fmt.Println("Bad result")
	}
}
