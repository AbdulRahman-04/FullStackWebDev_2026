package main

import (
	"fmt"
	"sync"
	"time"
)

func w1(ch1 chan string, wg *sync.WaitGroup) {
    defer wg.Done()
	time.Sleep(2*time.Second)
	ch1 <- "kaisa h bhai"

}

func w2(ch2 chan bool, wg *sync.WaitGroup) {
 
	defer wg.Done()
	time.Sleep(2500*time.Millisecond)
	// ch2 <- true
	ch2 <- false
	ch2 <- true
}

func main() {

	var wg sync.WaitGroup

	ch1 := make(chan string)
	ch2 := make(chan bool, 3)

	wg.Add(2)
	go w1(ch1, &wg)
	go w2(ch2, &wg)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg1 := <-ch2:
        fmt.Println(msg1)
	case msg2 := <- ch2:
		fmt.Println(msg2)
	case msg3 := <- ch2:
		fmt.Println(msg3)		
	}

	wg.Wait()
}
