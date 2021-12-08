// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.15: Types conversation
//
// @author Alex Versus 2021
//

package main

import (
	"fmt"
	"strconv"
)

// Entrypoint for program
func main() {
	var size, number int = 5, 12
	var meters float64 = 1.8

	//var length = number * meters // warning
	// converting
	var length = meters * float64(number) // 12.6

	fmt.Println(size, number, meters)
	fmt.Println(length)

	// example
	var widthOfHouse int = 60
	fmt.Println("Area of my house:", float64(widthOfHouse)*length)

	numOne, NumTwo := 5, 5.4
	fmt.Println("Compare one:", float64(numOne), NumTwo)
	fmt.Println("That more?", float64(numOne) < NumTwo)
	fmt.Println("Compare one:", numOne, int(NumTwo))
	fmt.Println("That more?", numOne < int(NumTwo))

	// Conversion to different types
	points := 130 // int
	bonus := 1.7  // float
	fmt.Println("Example about points:")
	fmt.Printf("%v, %T\n", points, points)
	fmt.Printf("%v, %T\n", bonus, bonus)
	result := int(bonus) * points
	fmt.Printf("%v, %T\n", result, result)   // 130 incorrect
	result2 := bonus * float64(points)       // 130.0 * 1.7
	fmt.Printf("%v, %T\n", result2, result2) // 221 float
	result3 := int(bonus * float64(points))
	fmt.Printf("%v, %T\n", result3, result3) // 221 int

	// Example with overflow
	fmt.Println("Overflow example:")
	var firstInt uint8 = 255
	var secondInt int = 255
	fmt.Printf("%v, %T\n", firstInt, firstInt)
	fmt.Printf("%v, %T\n", secondInt, secondInt)
	fmt.Println("uint8 == int : ", int(firstInt) == secondInt)
	secondInt++
	firstInt++
	fmt.Println("uint8 == int after ++: ", int(firstInt) == secondInt)

	// different size
	//var watermelon int = 5
	//var melon int64 = 10

	//melon = watermelon // warning
	//fmt.Printf("%v, %T\n", melon, melon)
	//fmt.Printf("%v, %T\n", watermelon, watermelon)

	// convert string
	var intExample = 9696
	fmt.Printf("%v, %T\n", intExample, intExample)
	fmt.Printf("Convert to string\n")
	fmt.Printf("%v, %T\n", string(intExample), string(intExample)) //◠, string
	var intNext = 66
	fmt.Printf("%v, %T\n", string(intNext), string(intNext)) //B, string

	fmt.Printf("%v\n", string([]byte{70, 75})) //in bytes

	// from string to float
	fmt.Printf("Convert from string to float \n")
	fromStr := "1000"
	toFloat, _ := strconv.ParseFloat(fromStr, 64)
	fmt.Printf("%f: %T\n", toFloat, toFloat)
	fmt.Printf("%g: %T\n", toFloat, toFloat)

	// conversion type in division
	fmt.Printf("Conversion to divider type\n")
	fmt.Printf("%v: %T\n", 4.5/2, 4.5/2)
	fmt.Printf("%v: %T\n", 3/1.5, 3/1.5)
	var floatT float64 = 6 / 5
	fmt.Printf("%v: %T\n", floatT, floatT) // 1 float
	floatT = float64(6) / 5
	fmt.Printf("%v: %T\n", floatT, floatT) // 1.2 float

	// correct convert
	converted := strconv.Itoa(intExample)
	fmt.Printf("%v, %T\n", converted, converted) //9696, string

	// warning
	//var strExample string = "5858"
	//fmt.Println("String to int?", int(strExample))

}
