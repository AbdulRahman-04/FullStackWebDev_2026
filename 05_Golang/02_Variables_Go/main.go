package main

import "fmt"

func main(){
  
	// fmt.Println("we seeing variables in go")

	// using var and const


	var age = 21
	fmt.Println(age)
	age = 22
    fmt.Println(age)

	// u cant change value of const variable
	const name = "rxhman"
	fmt.Println(name)

	// name = "heyy"
	// fmt.Println(name)

	var isAlive = true
	fmt.Println(isAlive)

	const isPresent = true
	fmt.Println(isPresent)


	// another way to decalre var in go is := which is most commonly used
	myClg := "dcet"
	fmt.Println(myClg)

	myFriend := "omer"
	fmt.Println(myFriend)

	myAge := 21
	fmt.Println(myAge)


	isAlive1 := true
	fmt.Println(isAlive1)
}