package main

import (
	aliasFmt "fmt"
	_ "image/png"
)

func main() {
	integers := []int{55, 55, 34}
	aliasFmt.Println(CalcSum(integers))

}

// Calc sum from slice of ints
func CalcSum(n []int) (s int) {
	for _, v := range n {
		s += v
	}
	return
}
