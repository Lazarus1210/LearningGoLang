package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const Max = 5000
	const NumSenders = 1000

	wgRecievers := sync.WaitGroup{}
	wgRecievers.Add(1)

	dataCh := make(chan int)
	stopCh := make(chan struct{})

	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	go func() {
		defer wgRecievers.Done()
		for value := range dataCh {
			if value == Max-1 {
				log.Printf("found value Max -1 : %d", value)
				close(stopCh)
				return
			}
			log.Println(value)
		}
	}()
	wgRecievers.Wait()
	log.Printf("Number of go routines %d", runtime.NumGoroutine())
}
