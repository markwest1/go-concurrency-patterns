# go-concurrency-patterns

To learn along with [Rob Pike's talk](https://www.youtube.com/watch?v=f6kdp27TYZs&t=708s).

Why do Golang's concurrency primitives promote creation of sophisticated, yet simple-to-reason-about concurrent programs, without the minutiae and memory barriers required by other languages?

* Locks
* Condition variables
* Callbacks

Watch Rob's talk to learn more.

Golang promotes composition of independently executable pieces of strait-forward sequential code to produce concurrent programs.

Some more examples to wrap your brain around:

* [Chat Roulette Toy](https://engineering.tumblr.com/post/23049359840/talk-by-andrew-gerrand-go-code-that-grows-with)
* [Concurrent Prime Sieve](https://play.golang.org/p/9U22NfrXeq)
* [Load Balancer](https://blog.golang.org/go-programming-session-video-from)
