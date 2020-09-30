package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		c <- 1
	}()
	fmt.Println("will wait for something to be added to this channel")
	<-c // main go routine blocked and waiting for some other go routine to send something on this channel
	fmt.Println("after value received: main")
}
