/**
Go (Golang) From simple to great. The Complete Developer's Guide.

Below is a program that emulates a customer database. When you query the database, a connection occurs, a search for viburnum.
If the client is found, the program displays information about it, if not, then displays an error-warning.
Rewrite the program so that if there is no client in the database, the connection to the database is closed anyway.


// Slice of company clients
type ClientList [] string

func main() {
	clientsDatabase := ClientList{"Mike", "Kelly", "Alex"}
	for _, client := range []string{"Alex","Alesya"} {
		err := clientsDatabase.FindClient(client)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Search clients to database
func (db ClientList) FindClient(name string) error {
	// Open connect to database
	db.Open()

	if find(name, db) {
		fmt.Println("Found", name) // fount client
	} else {
		return fmt.Errorf("%s not found", name)
	}
	db.Close()
	return nil
}

// Emulated Open connection to database
func (db ClientList) Open() {
	fmt.Println("Connection opened")
}

// Emulated Close connection to database
func (db ClientList) Close() {
	fmt.Println("Connection closed")
}
// Find clients in list
func find(title string, slice []string) bool {
	for _, sliceItem := range slice {
		if title == sliceItem {
			return true
		}
	}
	return false
}


@author Alex Versus 2021
www.alexversus.com
*/
package main

import (
	"fmt"
	"log"
)

// Slice of company clients
type ClientList []string

func main() {
	clientsDatabase := ClientList{"Mike", "Kelly", "Alex"}
	for _, client := range []string{"Alex", "Alesya"} {
		err := clientsDatabase.FindClient(client)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Search clients to database
func (db ClientList) FindClient(name string) error {
	// Open connect to database
	db.Open()

	defer db.Close()

	if find(name, db) {
		fmt.Println("Found", name) // fount client
	} else {
		return fmt.Errorf("%s not found", name)
	}

	return nil
}

// Emulated Open connection to database
func (db ClientList) Open() {
	fmt.Println("Connection opened")
}

// Emulated Close connection to database
func (db ClientList) Close() {
	fmt.Println("Connection closed")
}

// Find clients in list
func find(title string, slice []string) bool {
	for _, sliceItem := range slice {
		if title == sliceItem {
			return true
		}
	}
	return false
}
