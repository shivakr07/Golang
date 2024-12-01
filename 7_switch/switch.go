package main

import (
	"fmt"
	"time"
)

func main() {

	//simple switch
	i := 2

	switch i {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("one")

	case 3:
		fmt.Println("one")

	default:
		fmt.Println("other")

	}
	//we don't need to put the break like other languages

	//multi-condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("it's workday")
	}

	//type switch
	//if you keep the interface{} empty then i can hold any value
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		// switch i.(type) { -> if you don't want to use this value don't assign its type in t

		case int:
			fmt.Println("it's an integer value")

		case string:
			fmt.Println(("it's string"))

		case bool:
			fmt.Println(("it's a boolean value"))

		default:
			fmt.Println("it's others type", t)

		}
	}

	whoAmI(10)
	whoAmI("Hello")
}
