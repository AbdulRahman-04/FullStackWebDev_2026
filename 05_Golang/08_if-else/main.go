package main

import "fmt"

func main(){
	fmt.Println("enter a num to check even or odd")
	var num int

	fmt.Scan(&num)
	if num%2==0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}
}