// go routines are the lightweight threads
// when you want to run concurrent things or you want to do multithreading

package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing taks", id)
}

func main() {

	for i := 0; i <= 10; i++ {
		// task(i)    //blocking
		go task(i) //concurrent

		//an anonymous function
		//if you want to run this on a different goroutine then
		go func() {
			fmt.Println(i)
			// i is available due to closure
		}()

		//but it is a good practice that accept the variable in inline function
		//although closure is there but for good practice

		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
	//and also this is not a good practice to stop our main function
	//later we will see that how to do it better
}
