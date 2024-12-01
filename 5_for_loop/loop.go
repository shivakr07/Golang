package main

import "fmt"

//for -> only construct in go for looping

func main() {

	//while loop using for
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//while infinite loop
	// for {}

	//classic for loop
	fmt.Println("classic for loop")
	for j := 0; j < 5; j++ {
		if j == 2 {
			continue
			//it will skip for 2
		}

		if j == 4 {
			break
		}

		fmt.Println(j)

	}

	//range feature 1.22 go
	// runs for range-1
	for x := range 5 {
		fmt.Println(x)
	}

}
