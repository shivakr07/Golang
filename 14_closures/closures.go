package main

import "fmt"

// a function which returns a function whose return type is int
func counter() func() int {
	var count int = 0

	//we will return a anonymous function so that we can use count
	return func() int {
		count += 1
		return count
	}

}

func main() {

	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())

}

// we should get 1 1 ... because when you call a function then it loads in stack and once function ends it pops out of stack  so: count should be new everytime but
// due to closures we are able to get the same count and getting incremented value as 1 2 3 4 ...

// so bascially if you are using a variable inside a function of a function and it is using outer scope variable then it will always be availabe even after execution of the inner function
