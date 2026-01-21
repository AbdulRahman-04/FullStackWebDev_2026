package main

import (
	"fmt"
	"time"
)

func main(){
 
	// go m time pkg use krte current date and time nikalne...
	// go ka time format start hota Monday 02-Jan-2006 15:04:06
 
	// getDate := time.Now()
	// fmt.Println(getDate)  // isme time smjh nai ata toh best way h 

	// getDate := time.Now()
	// fmt.Println(getDate)

	formattedDate := time.Now().Format("Monday 02-jan-2006 15:04:06")
	fmt.Println(formattedDate)
	
}