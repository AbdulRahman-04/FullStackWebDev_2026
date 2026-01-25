package main

import "fmt"

type User struct {
	name string
	phone int
	isAlive bool
}

type Admin struct {
	name string
	age int
	isAlive bool
}

type myDetails struct {
	userName string
	age int
	isAlive bool
}


func main() {

	// Structs custom data types hote hain jo multiple values (alag-alag data types ki)
	// ko ek hi structure ke andar group karke store karte hain.

	//	Structs DB schema design ke liye use hote hain,
	// JSON se aane wala data structs me map karke store kiya ja sakta hai.
 
	// userDetails := User{
	// 	name: "rxhman",
	// 	phone: 5049,
	// 	isAlive: true,
	// }

	// fmt.Println(userDetails)

	// adminDetails := Admin {
	// 	name:  "fxhad",
	// 	age:  21,
	// 	isAlive: true,
	// }

	// fmt.Println(adminDetails)


	// myDetails1 := myDetails {
	// 	userName: "abdul rxhman",
	// 	age: 21,
	// 	isAlive: true,
	// }

	// fmt.Println(myDetails1)
}
