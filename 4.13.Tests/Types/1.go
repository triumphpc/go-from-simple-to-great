/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Write the conversion from centimeters to kilometers in the following example
--------
type (
	Meter      int64
	Centimeter int32
)

var (
	cm Centimeter = 1000
	m  Meter
)

m = Meter(cm) / 100
fmt.Println(m)
---------

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	type (
		Meter      int64
		Centimeter int32
		Kilometers float32
	)

	var (
		cm Centimeter = 1000
		m  Meter
		km Kilometers
	)

	m = Meter(cm) / 100
	fmt.Println(m)

	km = Kilometers(m) / 100
	fmt.Println(km)
}
