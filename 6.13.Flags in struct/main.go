//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.13: Tags in struct
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Example tag fot struct
type Example struct {
	Title  string   `json:"title"`
	Desc   string   `json:"desc"`
	Tags   []string `json:"tags,omitempty"`
	Secret string   `json:"-"`
}

// Entrypoint
func main() {
	e := &Example{
		Title:  "some title",
		Desc:   "some text this",
		Secret: "password",
	}

	fmt.Println(e)

	out, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))

}

func (e *Example) String() string {
	return fmt.Sprintf("Title is %s", e.Title)
}
