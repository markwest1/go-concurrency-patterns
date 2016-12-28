package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			cleanup()
			quit <- "See you!"
			return
		}
	}
}

func boring(msg string, q chan string) <-chan string {
	c := make(chan string)
	end := rand.Intn(10)

	go func() { // We launch the goroutine from inside the funtion.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			if i == end {
				q <- "Bye!"
				fmt.Printf("Joe says: %q\n", <-q)
			}
		}
	}()

	return c
}

func cleanup() {
	// Dummy function
}
