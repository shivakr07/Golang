package main

import "fmt"

//to send
// func processNum(numChan chan int) {

// 	// fmt.Println("processing number", <-numChan)

// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		//to make it little slow
// 		time.Sleep(time.Second)
// 	}

// }

// for receive
func sum(result chan int, num1 int, num2 int) {

	numResult := num1 + num2
	result <- numResult
}

func main() {

	// messageChan := make(chan string)

	// messageChan <- "ping"
	// msg := <-messageChan
	// fmt.Println(msg)
	// for send
	// -----------------------------
	//1:-  numChan := make(chan int)

	// go processNum(numChan)
	// // numChan <- 5 or
	// //generally we use it as a queue so this is how you can do it
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//----------------------------------
	// here we are using infinite loop so no need to use sleep to see the output as infinite loop is never ending goroutine
	// time.Sleep(time.Second * 2)
	//here we are using sleep because go processNum() is non blocking and  numChan <- 5 will run immediately and main will end

	//2:- for receving the data from the routine using channel

	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result
	//here we don't need sleep as it is blcking
	fmt.Println(res)

	// if you don't want to create a function outside and pass channel
	// go func(n1 int, n2 int) {
	// 	res := n1 + n2
	// 	result <- res
	// }(5, 5)
	// res := <-result
	// fmt.Println(res)

}
