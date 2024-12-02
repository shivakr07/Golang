package main

import "fmt"

func main() {

	//with slices
	nums := []int{6, 7, 8}
	for i, num := range nums {
		fmt.Println(i, "->", num)
	}
	//with maps
	m := map[string]string{"fname": "alan", "lname": "walker"}
	for k, v := range m {
		fmt.Println(k, "->", v)
	}
	//in maps if you use only sigle var then it gives you key
	for k := range m {
		fmt.Println(k)
	}

	//imp** will give you starting byte of rune [not index]
	//and ascii: [unicode code point rune] values
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
		//to get the value we need to convert it in string
	}
}
