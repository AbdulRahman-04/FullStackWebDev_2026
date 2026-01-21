package main

import "fmt"

func main(){
	// classic for loop : 
	// for i := 0 ; i<=3; i++ {
	// 	// fmt.Println(i)
	// }
 

	// while style for loop 
	// j := 1

	// for j<4 {
	// 	// fmt.Println(j)
	// 	// j++
	// }

	// break : breaks the loop when condition is satisfied.
	// for i := 5; i <= 15 ; i++{
		
	// 	if i == 10 {
	// 		break
	// 	}

	// 	fmt.Println(i)
		
	// }

	// for j := 2; j <= 20; j++ {
	// 	if j == 10 {
	// 		continue
	// 	}

	// 	fmt.Println(j)
	// }


	// for loop for looping on slices 

	makeSlice := make([]int, 0, 10)
	values := append(makeSlice, 12, 45, 23, 43, 55, 78)

	// for index, value := range values {
	// 	fmt.Printf("index is %d, value is %d\n", index, value)
	// }

	// for _ , values1 := range values {
	// 	fmt.Printf("values is %d\n", values1)
	// }

	for _ , values2 := range values {
		if values2%2==0 {
			fmt.Printf("valu is even %d\n", values2)
		}
	}
}