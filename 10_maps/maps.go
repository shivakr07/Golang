package main

import (
	"fmt"
	"maps"
)

// maps -> dictionary, hash
func main() {

	//create map
	m := make(map[string]string)

	//setting an element
	m["name"] = "golang"
	m["domain"] = "backend"

	//get an element
	fmt.Println(m["name"])

	//accessing a key which doesn't exist
	//it returns zeroed value : 0, false, "", ...
	fmt.Println(m["nokey"])

	//length of map
	fmt.Println(len(m))

	//to delete a key value from map
	delete(m, "domain")
	fmt.Println(m)

	//to clear the map
	clear(m)
	fmt.Println(m)

	//map without make

	dict := map[string]int{"price": 10, "discount": 5}
	fmt.Println(dict)

	v, ok := dict["price"]
	//multiple things get return in golang
	//v holds the value of key price and ok holds true of false [of existence]

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not okay")
	}
	fmt.Println(v)

	// to check whether they are equal or not we have package maps
	m1 := map[int]string{1: "Hello", 2: "Hii"}
	m2 := map[int]string{1: "Hello", 2: "Hii"}

	fmt.Println(maps.Equal(m1, m2))
}
