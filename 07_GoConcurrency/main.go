// package main

// import (
// 	"fmt"
// 	"sync"

// 	"time"
// )

// // assume apan image processing app bnare e.g resize, crop , filter
// func worker(url string, wg *sync.WaitGroup){

// 	defer wg.Done()
// 	time.Sleep(50 * time.Millisecond)

// 	fmt.Printf("image processsed %s\n", url)

// }

//   Fan-out = ek hi worker function ko multiple goroutines me concurrently chalana.

// func main(){
// 	// fmt.Println("Welcome to concurrency")

// 	// 1. use wait grp to track hamare go routines
// 	var wg sync.WaitGroup

// 	startTime := time.Now()

// 	wg.Add(5)
// 	go worker("image1.png", &wg)
// 	go worker("image2.png", &wg)
// 	go worker("image3.png", &wg)
// 	go worker("image4.png", &wg)
// 	go worker("image5.png", &wg)

// 	// fmt.Println(result1)
// 	// fmt.Println(result2)

// 	wg.Wait()

// 	fmt.Printf("it took %s ms\n", time.Since(startTime))

// }

// FAN IN :

package main

import (
	"fmt"
	"sync"

	"time"
)

type Result struct {
	Value string
	Err error
}

// assume apan image processing app bnare e.g resize, crop , filter
func worker(url string, wg *sync.WaitGroup, resultChan chan Result) {

	defer wg.Done()
	time.Sleep(50 * time.Millisecond)

	fmt.Printf("image processsed %s\n", url)

	resultChan <- Result{
		Value: url,
		Err: nil,
	}

}

func main() {
	// fmt.Println("Welcome to concurrency")

	// create channel for fan in : means worker func se data main go routine pe lana h
	resultChan := make(chan Result, 5)

	// 1. use wait grp to track hamare go routines
	var wg sync.WaitGroup

	startTime := time.Now()

	wg.Add(5)
	go worker("image1.png", &wg, resultChan)
	go worker("image2.png", &wg, resultChan)
	go worker("image3.png", &wg, resultChan)
	go worker("image4.png", &wg, resultChan)
	go worker("image5.png", &wg, resultChan)

	wg.Wait()
	close(resultChan)
	// reading data on main go routine,
	//  worker funcs jo light weight thread pe h
	// wo log iss main go routine channel pe write krre
	for result := range resultChan {
		fmt.Printf("recieved %v\n", result)
		if result.Err != nil {
			fmt.Println(result.Err)
		}
	}

	fmt.Printf("it took %s ms\n", time.Since(startTime))

}
