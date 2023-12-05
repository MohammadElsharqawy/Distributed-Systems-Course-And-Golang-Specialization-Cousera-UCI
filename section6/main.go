package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// chan <- int, channel of type sent-only
func sum(a int, b int) <-chan int { // return channel of type receive only

	return_channel := make(chan int, 1)

	go func() { // run in background
		delay := time.Duration(rnd.Int()%1000) * time.Millisecond
		time.Sleep(delay) // n7seso en 3malit el gm3 s3ba shwea...

		return_channel <- a + b
		close(return_channel)
	}()
	return return_channel
}

func print_ch(result int, ok bool, channel_name string) {
	// result: result of the summation.
	// ok, if the channel returns something.
	// name of the channel that i am receiving from.
	if ok {
		fmt.Printf("%v: 3 + 5 = %d\n", channel_name, result)
	}
	{
		fmt.Printf("%v: closed\n", channel_name)
	}
}

func main() {

	// call 4 times with the same parameters.

	// the delay of the sum is random, so all this function have different speeds
	// the first one finishes, the first one printed.
	ch1 := sum(3, 5)
	ch2 := sum(3, 5)
	ch3 := sum(3, 5)
	ch4 := sum(3, 5)

	for {
		select {
		case result, ok := <-ch1: // ok is false, when the channel is closed
			print_ch(result, ok, "channel 1")
		case result, ok := <-ch2:
			print_ch(result, ok, "channel 2")
		case result, ok := <-ch3:
			print_ch(result, ok, "channel 3")
		case result, ok := <-ch4:
			print_ch(result, ok, "channel 4")
		}
	}

	time.Sleep(1 * time.Second)
}
