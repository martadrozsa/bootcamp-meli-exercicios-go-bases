package main

import "fmt"

func main()  {
	var idade uint8 =  18
	switch {
	case idade >= 150:
		fmt.Println("É imortal?")
	case idade >= 18:
		fmt.Println("É maior de idade")
	default:
		fmt.Println("É menor de idade")
	}
}
