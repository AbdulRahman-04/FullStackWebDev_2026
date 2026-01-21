package main

import "fmt"

func main(){

	// creating array of fixed size length
	// arr := [3]string {"hi", "hey", "hello"}
	// fmt.Println(arr)

	// asking go to adjust to the size of array dynamically
	// myArr := [...]int {1,2,3,4,5,67}
	// fmt.Println(myArr[0], myArr[1], myArr[2])
	// fmt.Println(len(myArr))

	// arr1 := [2]string {"hello", "heyy"}
	// fmt.Println(arr1, len(arr1))
	// fmt.Println(arr1[0])

    dynamicArr := [...]int {1,2,3,4,5,5,6,7,7}
	fmt.Println(dynamicArr, len(dynamicArr))
}