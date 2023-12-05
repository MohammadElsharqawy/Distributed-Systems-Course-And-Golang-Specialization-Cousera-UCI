package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// it's going to send multiple random queries into the cache and the database.
	// Now, if the cache has a result, it's going to return that and print that out.
	// if the database has result, it's going to return it and print it out.
	// it is also going to update the cache to make sure that the in-memory cache has
	// that available for the next time we ask for it.

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1 // random num from 1 to 10
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("From Cache")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("From Database")
				fmt.Println(b)

			}
		}(id)
		//fmt.Printf("Book not found with id: '%v'\n", id)
		time.Sleep(150 * time.Millisecond)
		//to allow the cache time to receive its new values as we start to populate
		//that concurrently.

	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	// simulates that the database takes a longer time than cache
	for _, b := range Books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
