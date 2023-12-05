package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
