package main

import (
	"fmt"
	"io"
	"net/http"
)


func main(){
	url := "https://jsonplaceholder.typicode.com/posts/1"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer res.Body.Close()
	}

	// read data from res 
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}