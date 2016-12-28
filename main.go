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
	results := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
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
	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

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
