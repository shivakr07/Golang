package main

import "fmt"

const age = 30

//gender:= "male" --> you can't use the shorthand here

func main() {

	const name = "golang"
	fmt.Println(age)

	//you want to declare more than one constants at one time
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}
