package main

import (
	"fmt"
	"github.com/triumphpc/go-from-simple-to-great/go-from-simple-to-great/9.3.Createnewpackage/general"
)

func main() {
	integers := []int{55, 55, 34}
	fmt.Println(general.CalcSum(integers))

	house := general.House{Name: "New", Color: "Red"}
	fmt.Println(house)

	fmt.Println("It's type:" + general.BuildingType)

}
