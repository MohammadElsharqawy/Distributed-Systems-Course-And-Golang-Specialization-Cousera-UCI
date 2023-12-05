package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	fmt.Printf("New Routine\n")
	wg.Done()
}
func main() {
	// main func itself runs in a goroutine and it is created automatically
	//other go routine is crated by the keyword (go)
	/*
		here only one goroutine (main goroutine)
		a=1
		foo() main is blocked to execute the foo function
		a=2
	*/
	/*
		a=1 // main goroutine
		go foo() new goroutine, created a new thread executing while main doest get blocked
		a = 2 // this still executing while the new goroutine executing foo()
	*/
	// goroutine exit when its code is complete
	// when the main goroutine is complete, all other goroutines exit, so they might exit early before they are done with their job

	/*

		go fmt.Printf("New Routine\n")
		time.Sleep(1000 * time.Millisecond) // in main goroutine
		fmt.Printf("Main routine\n")

	*/

	// adding a delay is bad bad bad, bcz we are making assumptions about the time, and they may be wrong
	// may be the os takes all that time to do another operations (context switching)
	// may be go runtime schedules another goroutine
	// timing is non-deterministic

	// we need formal synchronization

	/////////////////////// basic synchronization ///////////////////////////
	// it is important to restrict ordering, restrict some of the different interleavings that are possible
	// os scheduler and go runtime schedule are deterministic algorithms but from our POV it is not deterministic bcz we dont understand their
	// algorithms
	// synchronization prevents bad interleavings (restrict the scheduling), reduce performance bcz sometimes the os will respect it
	// and be idel till it completes, but it is necessary in some cases (necessary evil lol)

	// wait groups: particular types of synchronization.
	// sync package
	// sync.WaitGroup ->>forces goroutine to wait for other goroutines, contains internal counter(counting semaphores)
	// increment it for each goroutine you wait for
	// decrement for each goroutine completes
	// cannot continue until the counter is 0
	// add() ->>> increments the counter
	// Done() ->>> decrements the counter
	// wait() ->>> blocks until the counter equals to 0

	var wg sync.WaitGroup
	wg.Add(1) // tell the main routine to wait for one goroutine
	go foo(&wg)
	wg.Wait() // make sure to write this to wait in the thread you want it to wait
	// we should make sure that each thread calls wg.Done() at the end when they are finished, we can use a defer

	fmt.Printf("Main Routine\n")

	////////////////////////////////////Communication/////////////////////////////

	////////////////////////////////synchronized communication/////////////////////

	c1 := make(chan int)
	c2 := make(chan int)

	for i := range c1 { // iteratively read from the channel, each time a channel receive a data, process it
		fmt.Println(i)
	} // this for loop will end when the sender close the channel (close(c))
	// we don't have to close the channel unless we use range keyword.

	// receive form multiple channels.
	// we can read them squentially
	//a := <- c1 // this is blocking, won't continue until receive the data from c1
	//b := <- c2 // and this as well
	//fmt.Println(a*b)

	//we can use select statement. we get from this one or this one or that one, first come, first served

	//select { // first come, first served and will skipp the second case
	//case a := <-c1:
	//	fmt.Println(a)
	//case b := <-c2:
	//	fmt.Println(b)
	//}

	// we may select to either send or receive operataions
	//b := 5
	//select { // which of them completes first, that is the one which executed
	//case a := <-c1:
	//	fmt.Println(a)
	//case c1 <- b:
	//	fmt.Println("sent b")
	//}

	//for { // keep receiving until an (abort signal) is received, abort may be a go routine
	//	select {
	//	case a := <-c1:
	//		fmt.Println(a)
	//	case <-abort:
	//		return
	//
	//	}
	//}
	// may want a default case to avoid blocking

	//select { // if we have a default we don't block, just go with it if neither of them is satisfied
	//case a := <-c1:
	//	fmt.Println(a)
	//case b := <-c2:
	//	fmt.Println(b)
	//default:
	//	fmt.Println("nop")
	//}

}
