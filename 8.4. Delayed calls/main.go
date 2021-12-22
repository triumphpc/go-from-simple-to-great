//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 8.4: Delayed calls, Defer
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
)

// Entrypoint
func main() {
	//defer fmt.Println("Goodbye!")
	//fmt.Println("Hello!")
	//fmt.Println("What?!")
	//
	//er := errFunc(second())
	//if er != nil {
	//	log.Fatal(er)
	//}

	//third example
	start()

	// arguments
	//nums := [...]int{0, 1, 2, 3, 4}
	//nums[0] += 1
	//defer fa(nums) // [1 1 2 3 4]
	//nums[0] += 5
	//fmt.Println(nums) // [6 1 2 3 4]

	//// slice example
	//slice := []int{0, 1, 2, 3, 4}
	//slice[0] += 5
	//defer fs(slice) // [10 1 2 3 4]
	//slice[0] += 5
	//
	// change outer func value
	fmt.Println(fchange()) // 4

	//// discarded value
	//fmt.Println(fdiscarted()) // 1
	//
	// anonymous
	_ = double(2) // double(2) = 4

}

func errFunc(i int) error {
	defer fmt.Println("one")
	fmt.Println("two")
	fmt.Println("five: ", i)

	return fmt.Errorf("error.")
}

func second() int {
	fmt.Println("four")
	return 1
}

func start() {
	fmt.Println("open")
	defer closeFunc()
	fmt.Println("finish")
}

func closeFunc() {
	fmt.Println("close")
}

func fa(nums [5]int) {
	fmt.Printf("%v\n", nums)
}

func fs(nums []int) {
	fmt.Printf("%v\n", nums)
}

func fchange() (v int) {
	defer func() {
		v *= 2
	}()
	v = 2
	return
}

func fdiscarted() int {
	defer func() int {
		return 10
	}()
	return 1
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}
