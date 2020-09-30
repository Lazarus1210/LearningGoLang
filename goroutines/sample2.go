package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Printf("go routine %d spwaned\n", i)
			<-c
			fmt.Printf("go routine %d after channel unblock\n", i)
		}(i)
	}
	fmt.Println("I will make the go routines wait for something to be added to this channel")
	time.Sleep(5 * time.Second)
	c <- 1                                    // main go routine blocked and waiting for some other go routine to send something on this channel
	fmt.Println("after value received: main") // the above send on the channel will just unblock 1 go routine and exit
	time.Sleep(2 * time.Second)
	fmt.Printf("num of go routines : %d\n", runtime.NumGoroutine())
}

/* output */
/*
coding@codingplatform:~/GoWorkSpace/src/LearningGoLang/goroutines$ ./sample2
I will make the go routines wait for something to be added to this channel
go routine 4 spwaned
go routine 2 spwaned
go routine 3 spwaned
go routine 0 spwaned
go routine 1 spwaned
after value received: main
num of go routines : 6

This is after putting a sleep of 2 seconds after sending value to the channel

coding@codingplatform:~/GoWorkSpace/src/LearningGoLang/goroutines$ ./sample2
I will make the go routines wait for something to be added to this channel
go routine 1 spwaned
go routine 2 spwaned
go routine 0 spwaned
go routine 4 spwaned
go routine 3 spwaned
after value received: main
go routine 1 after channel unblock
num of go routines : 5

Anyone of the goroutine can be unblocked after the value is seen in the channel
coding@codingplatform:~/GoWorkSpace/src/LearningGoLang/goroutines$ ./sample2
go routine 0 spwaned
go routine 1 spwaned
go routine 4 spwaned
I will make the go routines wait for something to be added to this channel
go routine 3 spwaned
go routine 2 spwaned
after value received: main
go routine 0 after channel unblock
num of go routines : 5


*/
