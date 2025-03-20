package main

import (
	"fmt"
	"maps"
)

func main() {

	// creating map

	// m := make(map[string]string)

	// setting an element in map

	// m["key1"] = "value1"
	// m["key2"] = "value2"

	// getting an element from map

	// fmt.Println(m["key1"])

	// IMPORTANT : if key does not exist it will return empty string

	// fmt.Println(m["key3"])

	// m := make(map[string]int)

	// m["key1"] = 10
	// m["key2"] = 20

	// getting an element from map
	// if key does not exist it will return 0

	// deleting an element from map
	// delete(m, "key1")

	// clearing the map
	// clear(m)

	// m := map[string]int{"key1": 10, "key2": 20}

	// ok := m["key1"] == 40

	// _, ok := m["key1"]

	// v, ok := m["key1"]

	// fmt.Println(v)

	// if ok {
	// 	fmt.Println("key1 exists")
	// } else {
	// 	fmt.Println("key1 does not exist")
	// }

	// fmt.Println(m)

	m1 := map[string]int{"key1": 10, "key2": 20}
	m2 := map[string]int{"key1": 10, "key2": 20}

	fmt.Println(maps.Equal(m1, m2))

}
