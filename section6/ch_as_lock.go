package main

import (
	"fmt"
	"time"
)

var signalling = make(chan struct{}, 1)

func worker1() {
	signalling <- *new(struct{})
	fmt.Println("using the shared resource1")
	time.Sleep(1 * time.Second)
	fmt.Println("done1")

	<-signalling

	time.Sleep(1 * time.Second)

}

func worker2() {
	signalling <- *new(struct{})
	fmt.Println("using the shared resource2")
	time.Sleep(1 * time.Second)
	fmt.Println("done2")

	<-signalling

	time.Sleep(1 * time.Second)
}

func main() {
	go worker1()
	go worker2()

	time.Sleep(3 * time.Second)
}
