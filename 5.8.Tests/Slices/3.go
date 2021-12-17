/**
Go (Golang) From simple to great. The Complete Developer's Guide.

Compare two slices

FOR EXAMPLE:
s1 := []byte{1,2,3}
s2 := []byte{2,1,3}

@author Alex Versus 2021
www.alexversus.com
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	s1 := []byte{1, 2, 3}
	s2 := []byte{2, 1, 3}
	s3 := []byte{1, 2, 3}

	res := bytes.Equal(s1, s2)
	fmt.Println(res) // false

	res = bytes.Equal(s1, s3)
	fmt.Println(res) // true

}
