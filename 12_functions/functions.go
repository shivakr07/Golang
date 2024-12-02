package main

import "fmt"

// if you don't tell returning type then you can't return
// you don't need to put the type with each arg if they are same type
// like a int, b int

func add(a, b int) int {
	return a + b
}

// you can return more than one value in go
// if they are more than then you have to tell type of each returned values

func getLanguages() (string, string, string, int) {
	return "golang", "js", "C", 10
}

// accept function
func processIt(fn func(a int) int) {
	fn(1)
}

// return function
func processingIt() func(a int) int {
	return func(a int) int {
		return 2
	}
}

func main() {

	result := add(3, 5)
	fmt.Println(result)

	fmt.Println(getLanguages())

	//to destructure kindof thing
	//** in project there are two values we return value and error
	lang1, _, lang3, num1 := getLanguages()
	fmt.Println(lang1, lang3, num1)
	//if you don't want to use anything use _

	//passing a function into function
	//anonymous function
	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	fn2 := processingIt()
	fn2(6)
}
