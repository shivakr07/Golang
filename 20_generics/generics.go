package main

import "fmt"

// introduced from go 1.18

// to print the items of int slice
// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// for string slice
// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// here you notice for each datatype we need to define a different function
// so there is duplication happening while inner logic is same
// we use generics

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// * you can use any or interface {} both are same
// but using any is not a good practice wo we can predefine what types are allowed
func printSlice2[T int | string | bool | float32](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// instead of defining the types using | we can use comparables in case of lot of types
// comparable is an interface which is implemented by all comparable types
func printSlice3[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// if you want to pass multiple types then also you can do that
func printSlice4[T comparable, V string](items []T, name V) {
	fmt.Println(items, name)
}

// -----------------------------------------
// till now we have seen generics with function only
// generics with structs
// this stack allow you to store elements as int
type stack struct {
	elements []int
}

// we use generics to store any type
type stack2[T any] struct {
	elements []T
}

// -----------------------------------------------
func main() {

	nums := []int{1, 2, 3, 4}
	printSlice(nums)

	names := []string{"golang", "typescript", "node"}
	// printStringSlice(names)
	printSlice(names)
	printSlice2(names)
	printSlice3(names)

	vals := []bool{false, true, false}
	printSlice4(vals, "alan")

	//generics with struct
	// ------------------------------------
	myStack := stack{
		elements: []int{1, 2, 3, 4},
	}
	fmt.Println(myStack)

	myStack2 := stack2[string]{
		elements: []string{"abc", "pqr", "wxy"},
	}
	fmt.Println(myStack2)
}
