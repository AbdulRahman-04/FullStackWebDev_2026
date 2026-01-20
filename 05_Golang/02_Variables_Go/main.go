package main

import "fmt"

func main(){
  
	// fmt.Println("we seeing variables in go")

	// using var and const


	// var age = 21
	// fmt.Println(age)
	// age = 22
    // fmt.Println(age)

	// // u cant change value of const variable
	// const name = "rxhman"
	// fmt.Println(name)

	// // name = "heyy"
	// // fmt.Println(name)

	// var isAlive = true
	// fmt.Println(isAlive)

	// const isPresent = true
	// fmt.Println(isPresent)


	// // another way to decalre var in go is := which is most commonly used
	// myClg := "dcet"
	// fmt.Println(myClg)

	// myFriend := "omer"
	// fmt.Println(myFriend)

	// myAge := 21
	// fmt.Println(myAge)


	// isAlive1 := true
	// fmt.Println(isAlive1)

	// println and printf
	// println prints in one line only and next print value is shifted to next line 
	name := "rxhman"
	age := 21
	isAlive := true
	clg := "dcet"
	marks := 66.54

	//println
	// fmt.Println("name:", name)
	// fmt.Println("age:", age)
	// fmt.Println("isAlive:", isAlive)
	// fmt.Println("clg:", clg)
	// fmt.Println("marks:", marks)

	//printf : iney format specifier smjhta and poori line nai leta iska baad ka print iske bazu hi print hota terminal mai. next line m print krwana h toh \n use kro.
	fmt.Printf("name is %s\n", name)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("is alive %t\n", isAlive)
	fmt.Printf("clg is %s\n", clg)
	fmt.Printf("marks is %f\n", marks)

	// format specifiers: %s : string, %d : integers, %t : boolean, %f: float values
    
}