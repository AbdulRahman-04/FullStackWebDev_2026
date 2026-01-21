package main

import (
	"errors"
	"fmt"
)

func AllDone(){
	fmt.Println("all functions executed")
}

func sum(x int, y int) (int, error) {
	if x%2!=0 || y%2!=0 {
		return  0 , errors.New("one of the numbers is odd")
	}

	return  x+y, nil
}

func main(){

	// defer jo h function k last m execute krta e.g idhr all done func sbse last execute hota kyuke ye first defer h 2nd defer 2nd last execute hota 3rd defer third last execute hota....

	defer AllDone()

	result, err := sum(14, 18)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	 defer fmt.Println("addition also done")
	
}