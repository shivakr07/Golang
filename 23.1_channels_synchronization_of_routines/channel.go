package main

import "fmt"

// we will achieve the wg functionality using channel
// wg to hold main until goroutines finish execution

// goroutine synchronizer
func task(done chan bool) {
	// fmt.Println("processing...")
	// done <- true
	//but this is not always right because if we get some error before we will never touch this
	//or we can use defer function [cleaning functions]
	//because after exiting the function it'll get surely

	defer func() {
		done <- true
	}()
	fmt.Println("processing...")

}

func main() {

	done := make(chan bool)

	go task(done)

	<-done //block until you get data from receiver
}
