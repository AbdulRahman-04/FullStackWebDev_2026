// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker1(wg *sync.WaitGroup, myChan chan int) {
// 	defer wg.Done()

// 	// sender ready
// 	myChan <- 160321735049
// }

// func main() {

// 	fmt.Println("main go routine started✅")

// 	var wg sync.WaitGroup

// 	// unbuffered channel: sender and receiver must be ready at the same time,
// 	// otherwise the goroutine will block and may cause a deadlock

// 	// creating unbuffered channel
// 	myChan := make(chan int)

// 	wg.Add(1)

// 	go worker1(&wg, myChan)

// 	// reciever ready
// 	msg := <-myChan
// 	fmt.Println(msg)

// 	wg.Wait()

// }

package main

import (
	"fmt"
	"sync"
)

func worker2(wg*sync.WaitGroup, buffChan chan int){
 defer wg.Done()
 buffChan <- 49
 buffChan <- 46
 buffChan <- 42
 buffChan <- 43
 buffChan <- 58
 buffChan <- 4044
 buffChan <- 4048
 buffChan <- 1604
 buffChan <- 1603
 buffChan <- 1602
}

func main(){
	fmt.Println("main go routine started✅")
 
	var wg sync.WaitGroup

	wg.Add(1)

// 	A buffered channel lets you send a fixed number of values without waiting for a receiver.
// Sending blocks only when the buffer is full, and receiving blocks only when the buffer is empty.

	// create buffered channel
	buffChan := make(chan int, 5)

	go worker2(&wg, buffChan)

	// reciver
	fmt.Println(<-buffChan, <-buffChan, <- buffChan, <- buffChan, <- buffChan)





	wg.Wait()
}