package main

import (
	"fmt"
	"time"
)

/*
A common requirement in programs is getting the number of seconds,
milliseconds, or nanoseconds since the Unix epoch. Here’s how to do it in Go.
*/

// Epoch func to illustrate Unix Epoch timing
func Epoch() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// You can also convert integer seconds or nanoseconds since
	// the epoch into the corresponding time.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
