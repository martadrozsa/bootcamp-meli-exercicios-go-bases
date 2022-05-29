package main

import (
	"fmt"
	"os"
)

func main()  {
	files, err := os.ReadDir(".")
	fmt.Println(files, err)
}
