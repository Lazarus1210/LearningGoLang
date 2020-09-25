package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// channel to send square of integers
var c = make(chan int)

// send square of numbers
func square(ctx context.Context) {
	i := 0

	for {
		select {
		case <-ctx.Done():
			return // kill goroutine
		case c <- i * i:
			i++
		}
	}
}

// main goroutine
/*func main() {

	// create cancellable context
	//An empty context is a Context that has no value, no deadline and it’s never canceled. The context.Background() function returns a default empty Context. This Context is generally used to derive other context objects since it never cancels. It can also be used in test cases or merely to pass a context object to an API where custom context is not important.

	ctx, cancel := context.WithCancel(context.Background())

	go square(ctx) // start square goroutine

	// get 5 square
	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	// cancel context
	cancel() // instead of `defer context()`

	// do other job
	time.Sleep(3 * time.Second)

	// print active goroutines
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}*/

