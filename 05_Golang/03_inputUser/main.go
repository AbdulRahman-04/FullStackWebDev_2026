package main

import (
	// "bufio"
	// "bufio"
	"fmt"
	// "os"
	// "os"
)

func main(){
	// take input from user, in go are two types : fmt.scan & bufio

	// fmt.scan:  cant read inputs after entering space while writing input.

	// fmt.Println("enter ur name")
	// var name string

	// fmt.Scan(&name)
	// fmt.Println("ur name is", name)


	// // bufio

	// fmt.Println("enter ur full name")
	// reader := bufio.NewReader(os.Stdin)
	// userIput, _ := reader.ReadString('.')
	// fmt.Println(userIput)

	// practice 

	// fmt.Println("enter ur details")

	// reader := bufio.NewReader(os.Stdin)

	// userInp, _ := reader.ReadString('.')
	// fmt.Println("user details are: ", userInp)


	fmt.Println("enter age")
	var age int

	fmt.Scan(&age)
	fmt.Println(age)


}