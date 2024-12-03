package main

import "fmt"

func sum(nums ...int) int {
	//nums here is slice of int
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// if you want to receive any type then you can use interface
func sum2(nums ...interface{}) {
	// total := 0
	// for _, num := range nums{
	// total += total + num
	//you can't do this because you don't know what kind of data you will be receiving due to interface
	// }
}

func main() {

	// fmt.Println(1, 2, 3, 4, 5, "hello")

	result := sum(3, 4, 5, 6, 7, 8)
	fmt.Println(result)

	//suppose if these numbers are in some slice then how to pass that ??
	vals := []int{3, 4, 5, 6}
	result2 := sum(vals...)
	fmt.Println(result2)

	//slice will be spread automatically in the function
}

/**
-> variadic function :
-> can receive n args using ...
-> ex :
	fmt.Println(1,2,3,4,5,"hello")
	-> this is receiving any type
	-> but we can restrict it to one type
	-> and you get slice of the args ex : nums above
-> you can pass n number of variables/ values in these functions
*/
