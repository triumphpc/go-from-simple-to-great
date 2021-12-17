//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 5.4: Maps
//
// @author Alex Versus 2021
/// www.alexversus.com

package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Entrypoint
func main() {
	// declaration
	var newMap map[string]int64
	newMap = make(map[string]int64)

	// short declaration
	secondMap := make(map[int]string)

	fmt.Println(newMap, secondMap)

	newMap["one"] = 555
	newMap["two"] = 666
	newMap["xxx"] = 7

	fmt.Println(newMap) // map[one:555 two:666 xxx:7]

	var l1, l2, l3 int = 5, 6, 7
	thirdMap := map[int]string{l1: "hi", l2: "buy", l3: "hello"}
	fmt.Println(thirdMap) // map[5:hi 6:buy 7:hello]

	// the same
	compare := make(map[int]string)
	compare[5] = "hi"
	compare[6] = "buy"
	compare[7] = "hello"

	// multilines
	el := map[int]int{
		1: 2,
		3: 4,
	}
	fmt.Println(el)

	// empty map
	emp := map[string]int{}
	fmt.Println(emp) // map[]

	// delete item of map
	delete(newMap, "two")
	fmt.Println(newMap)

	// add new item
	newMap["two"] = newMap["two"] + 10
	fmt.Println(newMap) // map[one:555 two:10 xxx:7]

	// loop and maps
	for st, in := range newMap {
		fmt.Println(st, in)
	}

	// sorting
	cars := map[string]string{
		"bmw":    "Germany",
		"audi":   "Germany",
		"ford":   "USA",
		"suzuki": "Japan",
		"lada":   "Russia",
	}

	all := make([]string, len(cars))
	for brand := range cars {
		all = append(all, brand)
	}
	fmt.Println("Before sort", all)
	sort.Strings(all)

	fmt.Println("After sort", all)
	for _, brand := range all {
		fmt.Printf("%s\t%s\n", brand, cars[brand])
	}

	// nil value error
	//var new map[string]int
	//new["test"] = 124

	// check key exist
	var addMap map[int]int
	addMap = make(map[int]int)
	addMap[1] = 2

	_, ok := addMap[2]
	if !ok {
		fmt.Println("addMap has no ")
	}

	// compare maps
	map1 := map[int]string{
		400: "A",
		401: "B",
		403: "C",
	}
	map2 := map[int]string{
		400: "A",
		401: "B",
		403: "C",
		404: "D",
	}

	map3 := map[int]string{
		400: "A",
		401: "B",
		403: "C",
	}
	map4 := map[string]int{
		"A": 400,
		"B": 401,
		"C": 403,
	}

	// Comparing maps
	// Using DeepEqual() function
	fmt.Println(reflect.DeepEqual(map1, map2)) // false
	fmt.Println(reflect.DeepEqual(map1, map3)) // true
	fmt.Println(reflect.DeepEqual(map1, map4)) // false
	fmt.Println(reflect.DeepEqual(map2, map3)) // false
	fmt.Println(reflect.DeepEqual(map2, map4)) // false
	fmt.Println(reflect.DeepEqual(map3, map4)) // false
	fmt.Println(reflect.DeepEqual(map4, map4)) // true
	fmt.Println(reflect.DeepEqual(map2, map4)) // false

	// composite type of item value in map
	var mapComposite map[string]map[string]int
	mapComposite = make(map[string]map[string]int)
	fmt.Println(mapComposite) // nil

}
