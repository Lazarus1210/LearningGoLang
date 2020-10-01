package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	buffersize := 5
	wordCh := make(chan string, buffersize)
	done := make(chan bool)
	// create a word array by splitting at the space delimiter
	words := strings.Split("my very educated mother show me the nine planets", " ")

	var wg sync.WaitGroup

	// spawning go routines
	// the default iterator in for range loop is always an integer and since the loop closure finishes before the go routines are spwaned the word pointer always points to the last word in the array.
	for _, word := range words {
		// Imp : in this buffered channel, the first 5 go routines will be spawned, they will send a word to channel and they will exit. But the remaiing 4 go routines for the remaining 4 words will be blocked at the channel, becasue channel is full to it capcaity, only when space is created then the go routines can write and finish
		wg.Add(1)
		go func(wordCh chan string, wg *sync.WaitGroup, someWord string) {
			time.Sleep(2 * time.Second)
			wordCh <- someWord
			wg.Done()
		}(wordCh, &wg, word)
	}

	// IMP this go routine with for range closure will forever be blocked waiting to read from this channel untill this channel is closed by the sender, thus channel closing principle becomes important. Thhe sender or a moderator must close this channel after it has sent all the words
	//wg.Add(1)
	// i will get out of the waitgroup man...because wait() is too fucked up
	// moral of the story : the main purpose of this go routine.
	/*
		1. recieve words using for range closure
		2. once i am done i will signal the main go routine and close the done channel
		3. somebody should close the for range channel on which i am blocked
	*/
	go func(d chan bool /*, wg *sync.WaitGroup*/) {
		for res := range wordCh {
			fmt.Print(fmt.Sprintf("%s\n", res))
		}
		fmt.Println("response1")
		close(d)
		fmt.Println("reponse 2")
		//wg.Done()
	}(done /*, /*&wg*/)

	/*go func() {
		fmt.Println("moderator go routine")
		//time.Sleep(10 * time.Second)
		done <- true
	}()*/
	// this will wait only this the time all the go routines which are part of waitgroup have returned. this will not care for any other go routine.
	wg.Wait()
	fmt.Println("after wait")
	//time.Sleep(time.Second * 2)
	close(wordCh)
	<-done
	fmt.Printf("number of go routines %d\n", runtime.NumGoroutine())
}
