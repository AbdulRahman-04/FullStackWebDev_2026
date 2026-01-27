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

func worker(jobs chan string, wg *sync.WaitGroup, bufChan chan string) {

	defer wg.Done()
	for url := range jobs {
		time.Sleep(2 * time.Second)
		bufChan <- url
	}

}

func main() {

	var wg sync.WaitGroup
	bufChan := make(chan string, 5)
	jobsChan := make(chan string)

	images := []string{
		"img1.png",
		"img2.png",
		"img3.png",
		"img4.png",
		"img5.png",
	}

	wg.Add(3)
	go worker(jobsChan, &wg, bufChan)
	go worker(jobsChan, &wg, bufChan)
	go worker(jobsChan, &wg, bufChan)

	for _, dataInsert := range images {
		jobsChan <- dataInsert
	}

	close(jobsChan)

	go func() {
		wg.Wait()
		close(bufChan)
	}()

	for data := range bufChan {
		fmt.Println(data)
	}
	fmt.Println("main go routine executed✅")

}
