// package main

// import (
// 	"fmt"
// 	"time"
// )

// func Add(x int, y int)  {

// 	fmt.Println(x+y)
// }

// func strPrint(str string) {

// 	fmt.Println(str)
// }

// func main() {

// 	fmt.Println("Main function executing on main thread")

// 	// 	A goroutine is a function that runs at the same time as other functions.

// 	// In Go, when you write go before a function:

// 	// that function runs in background

// 	// program does not wait for it to finish

// 	// Concurrency handles many tasks at once.
// 	// Parallelism runs many tasks at the same time.

// 	go Add(20,55)
//     go strPrint("hello")

// 	time.Sleep(7*time.Second)

// }

package main

import (
	"fmt"
	"time"
)

func task1(){
 
	time.Sleep(2*time.Second)

	fmt.Println("task1 done✅")

}
func task2(){
	time.Sleep(4*time.Second)

	fmt.Println("task2 done✅")
}

func task3(){
	time.Sleep(1*time.Second)

	fmt.Println("task3 done✅")
}


func main(){
	fmt.Println("main go routine started")

	go task1()
	go task2()
	go task3()

	time.Sleep(10*time.Second)
	fmt.Println("main go routine finished✅")
}