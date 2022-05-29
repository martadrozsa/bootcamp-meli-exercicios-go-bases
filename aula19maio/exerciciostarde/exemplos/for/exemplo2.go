package main

import "fmt"

/*
For range
*/

func main()  {

	frutas := []string{"maça", "banana", "morango"}

	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}
}