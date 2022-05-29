package main

import (
	"fmt"
	"os"
)

func main()  {
	data, err := os.ReadFile("./products.csv")
	fmt.Print(data, err)
}
