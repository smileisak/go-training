package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.

// Select func illustration Go's Select
func Select() {

	// For our example we’ll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time,
	// to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Two"
	}()

	// We’ll use select to await both of these values simultaneously,
	// printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		}
	}

	// Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

}
