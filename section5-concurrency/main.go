package main

import (
	"fmt"
	"time"
)

func wait_on_channel(ch chan int) {
	fmt.Println("Started Waiting.......")
	fmt.Println("received", <-ch)
}

func main() {
	// before defining the channel if we put <-, this channel is read-only.
	// we can't write into it.
	//ch := make(<- chan int)
	ch := make(chan int)

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")

		}
	}()

	go wait_on_channel(ch)
	time.Sleep(5000 * time.Millisecond)

	ch <- 5 //blocking <<<<<<<<<<<<>>>>>>>>>>>>>

	fmt.Println("hiiii")
	//value := <-ch  // it isnot gonna work, they should send & receive at the same time it is unbuffered
	//fmt.Println(value)

	time.Sleep(500 * time.Millisecond)

}
