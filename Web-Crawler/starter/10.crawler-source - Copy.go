package main

import (
	"fmt"
	"sync"
	"time"
)

var mut sync.Mutex
var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice. (visited_set has to be a thread-safe data structure)
	defer wg.Done()

	if depth <= 0 {
		return
	}

	//share memory by communication
	visited_set := <-ch
	visited, _ := visited_set[url]
	visited_set[url] = true
	ch <- visited_set

	if !visited {

		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found:[depth:%d] %s %q\n", depth, url, body)

		for _, u := range urls {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher)

		}

	}
}

func main() {
	wg.Add(1)

	go Crawl("http://golang.org/", 4, fetcher)
	wg.Wait()
	println("============DONE=============")

	for k := range visited_set {
		println(k)
	}

}

// visited_set
var visited_set = make(map[string]bool)
var ch = make(chan map[string]bool, 1)

// /////////////////////////////////////////////////////////////////////////////////////////////
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// ////////////////////////////////////////////////////////////////////////////////////////////
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

// every url has a body + urls
type fakeResult struct {
	body string
	urls []string
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {

	fmt.Printf("Fetching: %s\n", url)

	time.Sleep(500 * time.Millisecond)

	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

///////////////////////////////////////////////////////////////////////////////////////////////
