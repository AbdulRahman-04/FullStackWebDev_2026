package main

import (
	"errors"
	"fmt"
)

func multi(x int, y int) (int, error){
 
	if x %2!=0 || y%2!=0 {
		return  0, errors.New("error multiplying, one of the values is odd")
	}

	return x*y, nil
}

func main(){
	result, err := multi(20, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}