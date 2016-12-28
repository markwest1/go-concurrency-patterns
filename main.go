package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Search kinds
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

// Result is returned by Search for a query
type Result string

// Search receives a query and returns a result
type Search func(query string) Result

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

// First is called with a variadic set of search replicas to avoid placing a
// bet on a single query: only the results from the fastest search replica will
// be returned.
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

// Google a query
func Google(query string) (results []Result) {
	Web1 := fakeSearch("web1")
	Web2 := fakeSearch("web2")
	Image1 := fakeSearch("image1")
	Image2 := fakeSearch("image2")
	Video1 := fakeSearch("video1")
	Video2 := fakeSearch("video2")

	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
