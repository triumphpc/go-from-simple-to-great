//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// @author Alex Versus 2021
// www.alexversus.com
//
package main

func main() {
	for {
	s:
		switch 5 {
		case 5:
			switch 3 {
			case 3:
				break s
			}
		}
		break
	}
}
