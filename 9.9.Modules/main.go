package main

import (
	"fmt"
	m "github.com/triumphpc/go-fstg-modexam"
	m2 "github.com/triumphpc/go-fstg-modexam/v6"
)

func main() {
	fmt.Println(m.Hi("Alex"))

	fmt.Println(m2.Hi("Alex", "New string"))

}
