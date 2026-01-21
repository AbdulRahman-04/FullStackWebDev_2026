package main

import (
	"fmt"
	"strings"
)

func main(){
 
	// strings.Split string ko wahan todta hai
	//  jahan-jahan tumhara diya hua separator (jaise space " ") milta hai.
	
	str1 := "hell,obh,ai,kai,se,ho"
	words := strings.Split(str1, ",")
	fmt.Println(words)


	// 2️⃣ strings.Count – Count how many times a substring appears
	fruit := "strawberry"

	countR := strings.Count(fruit, "r")
	fmt.Println(countR)

	// 3️⃣ strings.TrimSpace – Remove leading/trailing spaces
	myHi := "          hye bro"
	spaceTirm := strings.TrimSpace(myHi)
	fmt.Println(spaceTirm)

	 // 4️⃣ strings.Join – Join a slice of words 
	 mySlice := make([]string, 0 ,3)
	 values := append(mySlice ,"rahman", "ismail", "omer", "saad")

	 joinVal := strings.Join(values, "")
	 fmt.Println(joinVal)

	  // 5️⃣ strings.HasPrefix – Check if string starts with prefix
	  file := "log.pdf.biodata.word"

	  checkWord := strings.HasPrefix(file, "log")
	  fmt.Println(checkWord)

	  myFile := "biodata.pdf"
	  checkFile := strings.HasSuffix(myFile, "pdf")
	  fmt.Println(checkFile)

}