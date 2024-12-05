package main

import (
	"fmt"
	"time"
)

// here we will achieve queue kind of functionaltiy using buffered channel
// in buffered channel we can send a limited amount of data without blocking
// like here for mailing
// in real life we will using struct instead of email string

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
		//simulating mail system otherwise will be very fast
	}
}

func main() {

	// emailChan := make(chan string, 100)

	// // it is non blocking until 100 data
	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	// -------------------------------------------------------
	// assume we are getting multiple emails from the db and we have send them in batch
	// implement queue system
	//done channel for blocking

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("test%d@gmail.com", i)
	}

	fmt.Println("done sending.")
	close(emailChan)
	<-done

}
