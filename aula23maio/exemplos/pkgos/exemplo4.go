package main

import (
	"fmt"
	"os"
)

func main()  {
	d1 := []byte("hello, gophers")
	err := os.WriteFile("./products.csv", d1, 0644)

	fmt.Println(err)
	fmt.Println(string(d1))
}