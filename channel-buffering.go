package main

import "fmt"

/*
By default channels are unbuffered,
meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready
to receive the sent value. Buffered channels accept a limited number of values without a corresponding
receiver for those values.
*/

// ChannelsBuffering Illustrate channel buffring in go
func ChannelsBuffering() {

	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered,
	//we can send these values into the channel without a corresponding concurrent receive.
	messages <- "Buffered"
	messages <- "Channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
