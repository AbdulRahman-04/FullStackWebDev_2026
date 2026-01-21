package main

import "fmt"

func simpleFunc(a int, b int) {

	fmt.Println(a+b)

}

// function with return
func funcReturn(x int, y int) int{
	return  x -y
}

// func with return
func mulFunc(i int, j int) int{
	return  i*j
}

//Variadic function wo hota hai jo ek hi parameter ke through unknown number of arguments accept kare, aur function ke andar wo parameter slice ki tarah treat hota hai.
func minus(numbers ...int)int{

	total := 0

	for _, num := range numbers {
     total = total + num
	}

	return  total
}

func variadicFunc(value ...int)int{
	total := 0

	for _, num := range value {
		total = total + num
	}

	return  total
}

func heavyMinus(mySlice ...int)int {

	total := 0

	for _, mySlice := range mySlice {
		total = total + mySlice
	}

	return total
}


func main(){
 
	simpleFunc(3,3)

	// return func handle 
	value := funcReturn(15,12)
	fmt.Println(value)

	value1 := mulFunc(3,10)
	fmt.Println(value1)

	// result := minus(34, 35, 56, 76, 100, 110)
	// fmt.Println(result)

	// result1 := variadicFunc(1,2,3,4,5,6,7,8,9)
	// fmt.Println(result1)

	result := heavyMinus(1,2,3,4,5,111,190, 122 )
	fmt.Println(result)

}