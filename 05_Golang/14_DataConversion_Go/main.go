package main

import (
	"fmt"
	"strconv"
)

func main(){

	// string to integar conversion
	myString := "5049"

	strToInt, err := strconv.Atoi(myString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("value is %d and type is %T\n", strToInt, strToInt)
	}


	// int to string 
    myInt := 4545

	intToStr := strconv.Itoa(myInt)
	fmt.Printf("Value is %s and datatype is %T\n", intToStr, intToStr)

	// string to float 
	flt1 := "12.34"

    strToFlt , err := strconv.ParseFloat(flt1, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("value is %f and type is %T\n", strToFlt, strToFlt)
	}


	// float to string 
	flt :=32.456

	fltToStr := strconv.FormatFloat(flt, 'f', 1, 64)
	fmt.Printf("value is %s and type is %T\n", fltToStr, fltToStr)
}