package general

import "fmt"

// Example of const
const BuildingType string = "house"

// Calc sum from slice of ints
func CalcSum(n []int) (s int) {
	for _, v := range n {
		s += v
	}
	return
}

type House struct {
	Name  string
	Color string
}

func (h House) String() string {
	return fmt.Sprintf("The house name is %q and is the color %s.", h.Name, h.Color)
}
