package main

import "fmt"

func main() {

	// age := 20
	// if age >= 18 {
	// 	fmt.Println("Person is adult", age)
	// } else {
	// 	fmt.Println("Person is minor")
	// }

	//else if
	umar := 20
	if umar >= 18 {
		fmt.Println("Person is Adult")
	} else if umar > 12 {
		fmt.Println("Person is teenger")
	} else {
		fmt.Println("Person is kid")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}

	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// we can declare the variable inside the if construct also
	// the best part of this syntax is you can use the variable inside the function, block and nested if else condition

	//note : go doesnot have ternary operator
	// we have to use normal if else for now 1.22
}
