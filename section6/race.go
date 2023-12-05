package main

import (
	"fmt"
	"sync"
	"time"
)

func race() {
	var mut = &sync.Mutex{}
	n := 0

	go func() {
		mut.Lock()
		n++
		fmt.Println("from anonymous function1", n)
		mut.Unlock()
	}()
	go func() {
		mut.Lock()
		n++
		fmt.Println("from anonymous function2", n)
		mut.Unlock()
	}()

}

func sharingByCommunication() {
	share_Ch := make(chan int)

	go func() {
		n := 0
		n++
		fmt.Println("From anonymous function", n)
		share_Ch <- n

	}()
	n := <-share_Ch
	n++
	fmt.Println("From race function", n)
}

func main() {
	//race()
	sharingByCommunication()
	time.Sleep(3000 * time.Millisecond)
}
