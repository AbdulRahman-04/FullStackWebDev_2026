package main

import "fmt"

func main(){
	// what is maps ?
	// maps basically ek unordered collection of data h jo key value pair m
	// data store krta and evry key unique rehti jisse uski value ko aap retrive kr skte
	// also map use krke aap data store krskte retrive krskte nd delete b kr skte

	// make function to create maps : 

	myMap := make(map[string] int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["hree"] = 3
	fmt.Println(myMap)

	// literal declarartions to create map 
    colors := map[int] bool{
		1 : true,
		0 : false,
	} 

	fmt.Println(colors)
	
}