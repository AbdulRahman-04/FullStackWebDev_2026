package main

import (
	"fmt"
	"net/url"
)

func  main()  {
	 
	myUrl := "https://github.com/AbdulRahman-04?tab=repositories"

	urlConv , err := url.Parse(myUrl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("type is %T\n", urlConv)
		fmt.Printf("scheme is %s\n", urlConv.Scheme)
		fmt.Printf("host is %s\n", urlConv.Host)
		fmt.Printf("path is %s\n", urlConv.Path)
		fmt.Printf("query is %s\n", urlConv.Query())
	}
	

}