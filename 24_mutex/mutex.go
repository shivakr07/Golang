package main

import (
	"fmt"
	"sync"
)

// mutex : mutual exclusion
// while we do multithreading to avoid race condition

// mimic post of socials
type post struct {
	views int
	mu    sync.Mutex
}

// when you call inc functin views get incremented
func (p *post) inc(wg *sync.WaitGroup) {
	// defer wg.Done()

	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	//since here modification is going on so
	p.mu.Lock()
	p.views += 1

}

func main() {

	myPost := post{views: 0} //default 0
	// myPost.inc()             //1
	// myPost.inc()             //2

	//but in real life things happens concurrently so it's not like let me first like the post then you

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}
