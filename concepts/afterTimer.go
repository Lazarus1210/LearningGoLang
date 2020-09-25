package main

import (
	"fmt"
)

var c chan Time

func handle(m int) {
	fmt.Println("inside handle \n")
	fmt.Println("%d", m)
}

/*
func main() {
	fmt.Println("time.Now()")
	<-time.After(5 * time.Second)
}*/
