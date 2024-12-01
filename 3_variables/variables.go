package main

import "fmt"

func main() {
	// var name string = "Evans"

	//auto infer the type if you don't put the assign the type
	var name = "Golang"
	fmt.Println(name)

	var isAdult = true
	fmt.Println(isAdult)

	var age int = 20
	fmt.Println("Age is ", age)

	//shorthand
	//we use this when we want to declare and assign value at the same time
	gender := "male"
	fmt.Println(gender)

	var career string
	career = "software engineer"
	fmt.Println(career)

	// if you are declaring the variable then you have to put the type and keyword both

}
