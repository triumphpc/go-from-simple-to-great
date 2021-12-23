package main

import (
	"fmt"
	"main/github.com/triumphpc/go-cource/9.3/general"
)

func main() {
	integers := []int{55, 55, 34}
	fmt.Println(general.CalcSum(integers))

	house := general.House{Name: "New", Color: "Red"}
	fmt.Println(house)

	fmt.Println("It's type:" + general.BuildingType)

}
