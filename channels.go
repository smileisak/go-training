package main

import "fmt"

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.
*/

// Channels illustrate channels in go
func Channels() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey (transmettre).
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() {
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel.
	// Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	// go func() {
	// 	m := <-messages
	// 	fmt.Println(m)
	// }()

	// When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
}
