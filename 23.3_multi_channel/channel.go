package main

import "fmt"

//we will listen/receive here from multiple channels

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from the chan2", chan2Val)
		}
	}
}
