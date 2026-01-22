package main

import (
	"context"
	"fmt"
	"time"
)

func sum(ctx context.Context, x int, y int) {

	for {
		if ctx.Err() != nil {
			fmt.Println("STOPPED❌")
			return
		}
		fmt.Println("sum is", x+y)
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	sum(ctx, 10, 10)
}

// The context package is used to stop ongoing work 
// when it is no longer needed, such as when a request times out or is cancelled.


// A user hits an API.
// You allow 10 seconds for processing.
// If it takes longer, you stop everything.

// What context does

// Starts a timer (10s)

// After 10s, sends a “stop” signal

// Your code sees the signal and exits