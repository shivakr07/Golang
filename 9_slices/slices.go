package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array
// most used construct in go
//+ useful methods

func main() {
	// uninitialized slice is nil in go [null]
	var nums []int
	// fmt.Println(nums == nil) // []
	fmt.Println(len(nums)) // 0

	// if you don't want it to be null
	// you need to pass three things type, initial size, initial capacity
	// it behaves like array but it is dynamic
	// putting initial size doesn't limits its size

	// capacity -> max no of elements can fit
	// if you don't put capacity it auto become size of initial size

	// adding capacity and length doesn't restrict its size its dynamic so size and len both are dynamic also
	// var number = make([]int, 0, 5)
	var number = make([]int, 2, 5)

	fmt.Println(number)
	fmt.Println(cap(number))

	number = append(number, 1)
	number = append(number, 2)
	number = append(number, 3)
	number = append(number, 4)

	// we can use index also to get and add the items
	number[0] = 10

	fmt.Println(number)
	fmt.Println(cap((number)))
	// [0 0 1 2 3 4]  -> 2 added later 2
	// so when items increases then its capacity then arr becomes of double as other memory resizes instead of 1 by 1 as padding happens
	// like capacity should be 6 here while it is 10 now
	// because initial length was 2 so initial things were [0 0]
	// when appended 1 2 3 and 4 then it became like this
	// [0 0 1 2 3 4]
	// now as 4 appended it become the 6th member while cap was 5 so it became
	// 5*2 -> 10

	//** if you notice we are getting 0 0 due to initial length of slice
	// to avoid that as getting 0s we keep initial length 0

	//more useful syntax -> similar to python list
	li := []int{}
	fmt.Println(li)

	li = append(li, 1)
	li = append(li, 2)

	//to copy the slice
	var nums1 = make([]int, 0, 5)
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 4)

	var nums2 = make([]int, len(nums1))

	copy(nums2, nums1)
	fmt.Println("dest slice ", nums2)
	fmt.Println("src slice ", nums1)

	// slice operator [similar as python]
	var item = []int{1, 2, 3}
	fmt.Println(item[0:2])
	fmt.Println(item[:2])
	fmt.Println(item[1:])

	//slices : package
	var prod1 = []int{1, 2}
	var prod2 = []int{1, 2}
	fmt.Println(slices.Equal(prod1, prod2))

	//2D slice
	var x = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(x)
}
