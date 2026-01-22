// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func w1(makeChan chan bool, wg *sync.WaitGroup){
//    defer wg.Done()
//    time.Sleep(2*time.Second)
// 	makeChan <- true

// }

// func main(){
// 	var wg sync.WaitGroup

// 	makeChan := make(chan bool)
//     wg.Add(1)
// 	go w1(makeChan, &wg)

// 	msg := <- makeChan
// 	fmt.Println(msg)
// 	fmt.Println("worker 1 finished✅")

// 	wg.Wait()
// 	fmt.Println("main go routine finished✅")
// }

package main

import (
	"fmt"
	"sync"
)

func worker2(makeChan chan int, wg*sync.WaitGroup){
  defer wg.Done()
	makeChan <- 200
	makeChan <- 400
	makeChan <- 500
	makeChan <- 501
	makeChan <- 502
	makeChan <- 600
}

func main(){
 
	var wg sync.WaitGroup

	makeChan := make(chan int, 5)
    wg.Add(1)
	go worker2(makeChan, &wg)

	fmt.Println(<- makeChan)
	fmt.Println(<-makeChan)
	fmt.Println(<- makeChan)
	fmt.Println(<- makeChan)
	fmt.Println(<-makeChan)
	fmt.Println(<- makeChan)

	wg.Wait()
	fmt.Println("main go rotuine finished")
}