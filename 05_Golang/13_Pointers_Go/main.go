// package main

// import "fmt"

// func main(){

// 	// pointers means ek variable k value ka memory address
// 	myVar := 12

// 	myPtr := &myVar   // & se variables ka memory address nikalte
// 	                  // * se variables k memory m value kya stored h woh aati

// 	fmt.Println(myPtr, *myPtr)

// 	modifyValueByRefernce(myPtr)
// 	fmt.Println(*myPtr)

// }

// func modifyValueByRefernce(myPtr *int){
// 	*myPtr = 160321735049
// }

package main

import "fmt"

func modifyValueByRefernce(myPtr *int){
	*myPtr = 1605
}

func main(){
	myVar := 3656

	myPtr := &myVar

	fmt.Println(myPtr, *myPtr)

	modifyValueByRefernce(myPtr)
	fmt.Println(*myPtr)

}