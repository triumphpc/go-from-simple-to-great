//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.8: Concat and conversions
//
//@author Alex Versus 2021
//
// www.alexversus.com
//

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Entrypoint
func main() {

	// Simple example
	fmt.Println("One" + "Two")

	// Error
	//fmt.Println("Summarize with int" + "1")

	//covert int to string
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

	// convert int to binary
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	// convert string to int
	z1, _ := strconv.Atoi("123")
	z2, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(z1, z2)

	// to uppercase
	st := "Example String"
	fmt.Println(strings.ToUpper(st))

	// to lowercase
	fmt.Println(strings.ToLower(st))

	// search string
	str := "Test string"
	fmt.Println(strings.HasPrefix(str, "Test"))   // true
	fmt.Println(strings.HasPrefix(str, "T"))      // true
	fmt.Println(strings.HasPrefix(str, "string")) // false

	fmt.Println(strings.HasSuffix(str, "Test"))   // false
	fmt.Println(strings.HasSuffix(str, "string")) // true
	fmt.Println(strings.HasSuffix(str, "ing"))    // true

	fmt.Println(strings.Contains(str, "st")) // true

	fmt.Println(strings.Count(str, "s")) // 2

	// compare
	fmt.Println("One" == "one") // false

	fmt.Println(strings.Count("Super start for startup", "s")) // 2
	fmt.Println(strings.Count("Super start for startup", "S")) // 1

	// get length
	fmt.Println(len("this is big string")) // 18

	// join strings from slices
	strs := []string{"one", "two", "three"}
	out := strings.Join(strs, ",")
	fmt.Println(out) // concatenates substrings from slices

	splitOut := strings.Split(out, ",")
	fmt.Printf("%T %[1]v - %[1]q\n", splitOut)

	// Field example
	fieldOut := strings.Fields("this is ,     some  interesting text")
	fmt.Printf("%T %[1]v - %[1]q\n", fieldOut)

	// replace some substring
	r := "Our life is what our thoughts make it. Avreliy"
	fmt.Println(strings.ReplaceAll(r, "Avreliy", "Alex Versus"))

}
