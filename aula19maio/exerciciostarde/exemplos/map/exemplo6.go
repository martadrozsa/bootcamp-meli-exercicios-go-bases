package main

import "fmt"

/*
Percorrendo elementos de um map
*/

func main()  {
	var students = map[string]int{"João": 20, "Pedro": 26}
	for key, element := range students {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}