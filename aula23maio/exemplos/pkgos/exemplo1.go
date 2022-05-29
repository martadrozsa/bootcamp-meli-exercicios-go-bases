package main

import (
	"fmt"
	"os"
)

func main()  {
	err := os.Setenv("Name", "Gopher")
	fmt.Println(err)

	value := os.Getenv("Name")
	fmt.Println(value)

	value, ok := os.LookupEnv("Name")
	fmt.Println(ok)

}

