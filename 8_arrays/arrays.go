package main

import "fmt"

// numbered sequence of specific length
func main() {

	var nums [4]int
	fmt.Println(len(nums))

	//to add and get an element in the array
	nums[0] = 10
	fmt.Println(nums[0])

	fmt.Println(nums)

	//if you don't assign the values then it becomes 0 for int type
	// if it is of another type then zeroed value for that type
	//like for bool -> false
	//int => false
	//strings -> ""   empty array [ ]
	var sts [2]string
	fmt.Println(sts)
	fmt.Println(sts[0])

	//assign the values while declaring the array [in single line]
	//you can't assign more values then length
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	//2-d arrays
	counts := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(counts[0][1])
	fmt.Println(counts)

	//fixed size, that is predicatable
	//memory optimization
	//const time access

	// note we use arrays when we know the size of the array
	// as it is memory efficient language
	// otherwise we use slices
	// generally we use slices as it is dynamic in nature don't have to thing about the size
}
