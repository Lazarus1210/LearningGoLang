package main

import (
	"fmt"
	"time"
)

// Client struct
type Client struct {
	ca int
}

// Mathematics struct
type Mathematics struct {
	a int
	b int
	c int
}

type stateFunc func(*Mathematics) int // this is nothing but a closure, that needs to be invoked. callbacks are more of event driven and they report bakc to the caller when the assign jobs are done. Callback are asynchronous in nature

func (c *Client) addition(m *Mathematics) int {
	fmt.Println("inside the callback Addition, add with clients value")
	return m.a + m.b + m.c
}

func (c *Client) run(m *Mathematics, f stateFunc) {
	fmt.Println("inside run")
	for {
		fmt.Println("inside run: inside for loop ")
		time.Sleep(5 * time.Second)
		a := f(m)
		fmt.Println("output from callback is %d", a)
	}

}

//commenting main...uncomment to run it...

/*
func main() {
	mainClient := &Client{4}
	math := &Mathematics{1, 2, 3}
	mainClient.run(math, mainClient.addition)

}*/
