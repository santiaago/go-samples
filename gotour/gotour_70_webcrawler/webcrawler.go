package main

import (
	"fmt"
	"errors"
	"sync"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

var inprogress = errors.New("in progress")

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	
	fetched.Lock()
	if _, ok := fetched.m[url]; ok{
		fetched.Unlock()
		return
	}
	fetched.m[url] = inprogress
	fetched.Unlock()

	body, urls, err := fetcher.Fetch(url)
	
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	finnished := make(chan bool)
	
	for _, u := range urls {
		go func(url string){
			Crawl(url, depth-1, fetcher)
			finnished <- true
		}(u)
	}
	for _ = range urls{
		<-finnished
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
	for u, err := range fetched.m{
		if err != nil{
			fmt.Printf("%s ko: %s",u,err)
		} else{
			fmt.Printf("%s ok",u)
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
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
