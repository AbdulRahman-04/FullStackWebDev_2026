// sync.WaitGroup is used to wait for a collection of goroutines to finish.
// It helps the main goroutine wait until all required goroutines complete their work.

// Why we need sync.WaitGroup:
// - Goroutines run independently and do NOT block main.
// - If main finishes early, the program exits and kills all goroutines.
// - WaitGroup allows us to control program exit by waiting for specific goroutines.

// How it works:
// 1. wg.Add(n) tells Go how many goroutines to wait for.
// 2. Each goroutine calls wg.Done() when its work is finished.
// 3. wg.Wait() blocks main until the counter becomes zero.

// Important:
// - WaitGroup only waits for goroutines that call Done().
// - If a goroutine is not added to the WaitGroup, it will NOT be waited for.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(wg *sync.WaitGroup){
	defer wg.Done()

	time.Sleep(2*time.Second)
	fmt.Println("worker 1 done🚀")

}

func worker2(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("worker 2 done🚀🚀")

}

func worker3(wg*sync.WaitGroup){
	defer wg.Done()
	fmt.Println("worker 3 done🚀🚀🚀")

}
func worker4(){
  time.Sleep(6*time.Second)
  fmt.Println("worker 4 done")
}


func main(){

	var wg sync.WaitGroup

	wg.Add(3)

	go worker1(&wg)
	go worker2(&wg)
	go worker3(&wg)
	worker4()

	wg.Wait()

}

// We pass the memory address of WaitGroup in all go routines so that ALL goroutines work on the SAME WaitGroup.

// If you pass the whole variable → each goroutine gets its OWN COPY ❌
// Then Done() will NOT affect the original one in main.