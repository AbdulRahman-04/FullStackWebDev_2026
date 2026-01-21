package main

import "fmt"

func main(){
 
	// slice? under the hood slice array hi h jo data hold krte apne andar
	// waise toh arr b use krskte but slices are more dynamic jisku ... operator dene ki zroorat nai arr ku dynamic bnane size ke according
	mySlice := []int {1,2,3,4}
	fmt.Println(mySlice)

	mySlice1 := []bool {true, false, true}
	fmt.Println(mySlice1)

	mySlice2 := []string {"hey", "kaise ho?", "rahman?"}
	fmt.Println(mySlice2)


	// u can use slice method of js array on  slice of go
	slice := mySlice2[0:1]
	fmt.Println(slice)
	slice2:= mySlice1[0:2]
	fmt.Println(slice2)
	slice3:=mySlice[0:1]
	fmt.Println(slice3)

	// using make function to create slices

	// capacity 5 h matlab bina new arr/slice bnaye iss slice m kitne elems aa skte agar capacity badhgyi toh append ek naya bigger size slice bnake sab elems usme copy krta 
	
	normalSlice := []int {1,2,3}
	fmt.Println(normalSlice)

	makeSlice := make([]bool, 0, 5)
	values := append(makeSlice, true, true, true ,false)
	fmt.Println(values)
}