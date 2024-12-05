// in the previous lesson you see sleep : here we have waitgroups
// because in real world we don't know how much time a task is going to take

//\\ wait groups : are a mechanism to synchronize our go routines //\\

package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	//after executing the function we have to remove 1 from wg added with Add
	defer w.Done()
	//what ever you write in defer that get executed after function execution

	fmt.Println("doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		//its counter type mechanism
		wg.Add(1)
		//here we are adding 1 because one go routine is starting

		go task(i, &wg)
	}

	wg.Wait()

	// time.Sleep(time.Second * 2)
	//we need this because until our goroutines run main is over
	/*
		add
		done
		wait
	*/

}
