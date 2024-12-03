package main

import "fmt"

// by value
func changeNum(num int) {
	num = 5
	fmt.Println("in changeNum1", num)
}

// by reference
func changeNum2(num *int) {
	*num = 5
	fmt.Println("in changeNum2", *num)
}

func main() {

	num := 1

	//pass by value
	changeNum(num)
	fmt.Println("After change num in main", num)

	//pointers
	num2 := 10
	// fmt.Println("Memory address", &num2)

	fmt.Println("Number 2 before change", num2)
	changeNum2(&num2)
	fmt.Println("Number 2 after change", num2)
}
