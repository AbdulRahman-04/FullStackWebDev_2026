package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
	isAlive bool
}

func main(){
 
	myDetails := Person {
		Name: "rxhmannnn",
		Age: 21,
		isAlive: true,
	}

	// convert the struct into json
	jsonData, err := json.Marshal(myDetails)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(jsonData))
	}

	// unmarshall json data into struct form 
	var unmarshall Person

	err = json.Unmarshal(jsonData, &unmarshall)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(unmarshall)
	}

}