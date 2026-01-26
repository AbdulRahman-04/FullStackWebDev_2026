// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(url string, wg*sync.WaitGroup, dataChan chan string){
//   defer wg.Done()
// 	time.Sleep(2000*time.Millisecond)
// 	fmt.Printf("url processed: %s\n", url)

// 	dataChan <- url
// }

// func main(){

// 	var wg sync.WaitGroup

// 	dataChan := make(chan string, 5)

// 	wg.Add(3)
// 	go worker("image1.png", &wg, dataChan)
// 	go worker("image2.png", &wg, dataChan)
// 	go worker("image3.png", &wg, dataChan)

// 	wg.Wait()
//     close(dataChan)

// 	for realData := range dataChan {
// 		fmt.Printf("data inside channel is: %s\n", realData)
// 	}

// 	fmt.Println("main go routine executed✅")

// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobs chan string, wg *sync.WaitGroup, mainChan chan string) {

	defer wg.Done()
	for url := range jobs {
		time.Sleep(1*time.Second)
		mainChan <-url
	}

}

func main() {

	var wg sync.WaitGroup
	// channel to recive data from worker thread
	mainChan := make(chan string, 10)
	// create jobs channel
	jobs := make(chan string)

	images := []string{
	"image1.png",
	"image2.png",
	"image3.png",
	"image4.png",
	"image5.png",
	"image6.png",
	"image7.png",
	"image8.png",
	"image9.png",
	"image10.png",
}


	wg.Add(3)
	go worker(jobs, &wg, mainChan)
	go worker(jobs, &wg, mainChan)
	go worker(jobs, &wg, mainChan)
 
	// slice k andar data -> jobs channel
	for _, img := range images {
		jobs <- img
	}

	close(jobs)

	go func ()  {
	wg.Wait()
	close(mainChan)	
	}()
 
	for data := range mainChan {
      fmt.Println("channel data is:", data)
	}

	fmt.Println("main ggo routine executed✅")
}
