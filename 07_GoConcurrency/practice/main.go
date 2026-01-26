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

func worker(jobs chan string, wg *sync.WaitGroup, mainChann chan string) {
	defer wg.Done()
	for url := range jobs {
		time.Sleep(1 * time.Second)
		fmt.Printf("url processed %s\n", url)

		mainChann <- url
	}
}

func main() {

	var wg sync.WaitGroup

	jobs := make(chan string)
	mainChann := make(chan string, 5)

	// go worker("image1.png", &wg, mainChann)
	// go worker("image2.png", &wg, mainChann)
	// go worker("image3.png", &wg, mainChann)

	// SLICE of jobs
	images := []string{
		"image1.png",
		"image2.png",
		"image3.png",
		"image4.png",
		"image5.png",
	}

	wg.Add(3)
	go worker(jobs, &wg, mainChann)
	go worker(jobs, &wg, mainChann)
	go worker(jobs, &wg, mainChann)

	for _, img := range images {
		jobs <- img
	}
	close(jobs)

	wg.Wait()
	close(mainChann)

	for date := range mainChann {
		fmt.Println(date)
	}
	fmt.Println("main go routine executed✅")

}
