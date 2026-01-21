package main

import "fmt"

func main() {
	day := 4

	switch day {
	case 0:
		fmt.Println("sunday")
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	default:
		fmt.Println("not a day")
	}

	temp := 33

	switch temp {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		fmt.Println("cold")
	case 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		fmt.Println("cool")
	case 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
		fmt.Println("a bit hot")
	case 31, 32, 33, 34, 35:
		fmt.Println("hot")
	default:
		fmt.Println("hottest weathet")
	}
}
