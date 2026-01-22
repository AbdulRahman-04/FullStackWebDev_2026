package main

import (
	"context"
	"fmt"
	"time"
)

func sum(ctx context.Context,x int, y int) {

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
