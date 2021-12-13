/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Define two defined types, for example Glass (200 ml) and Jar (1 L)
2. Write conversion functions
3. Convert values from one to another type and vice versa

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

type (
	Glass float32
	Jar   float32
)

func main() {
	var g Glass = 10
	fmt.Println("I have glasses", g)

	mlVolume := getMlFromGlasses(g)
	fmt.Println("I't mls ", mlVolume)

	fmt.Println("I't count of jars ", toJar(g))
}

// milling volume in a glass
func getMlFromGlasses(gc Glass) float32 {
	return float32(gc * 200)
}

func toJar(gc Glass) Jar {
	return Jar(getMlFromGlasses(gc) / 1000)
}
