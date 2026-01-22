package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, ch chan string){
	defer wg.Done()

	// sender ready
	ch <- "ye le bhai worker go routine ka data"

}

func main(){
 
	var wg sync.WaitGroup

	// create a chan
	ch := make(chan string)

	wg.Add(1)

	go worker(&wg, ch)

	// reciever ready
    chanMsg := <- ch

	fmt.Println(chanMsg)

	wg.Wait()

}