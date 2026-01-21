// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main(){

// 	// creating a file
//     fileCreate, err := os.Create("sendEmail.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("file created: ", fileCreate)
// 	}

// 	// write content inside file
// 	_, err = fileCreate.WriteString("paise pay kro\n kab ta phugat ki?")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("content added✅")
// 	}

// 	// open the file to read its content
//     openFile, err := os.Open(fileCreate.Name())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		defer openFile.Close()
// 	}

// 	// create a byte slice to read data from fiel in chunks

//     buffer := make([]byte, 1024)

// 	bytesRead, err := openFile.Read(buffer)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {

// 		fmt.Println("bytes read:", bytesRead)
// 		fmt.Println("file content:\n", string(buffer[:bytesRead]))

// 	}

// }

package main

import (
	"fmt"
	"os"
)


func main(){
	// create a file 
	fileCreate, err := os.Create("sendEmail.txt")
	if err != nil {
		fmt.Println(err)
		return
	} else {
       fmt.Println("file created✅")
	}

	// add content in file 
	_, err = fileCreate.WriteString("paise de bhai\n kab takh phugat ki")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("content added✅")
	}

	// open file to read 
	openFile, err := os.Open(fileCreate.Name())
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer openFile.Close()
	}

	// create buffer to read data in chunks
	buffer := make([]byte, 1024)

	byteRead, err := openFile.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("bytes read is:\n", byteRead)
		fmt.Println("file data is:", string(buffer[:byteRead]))
	}
}