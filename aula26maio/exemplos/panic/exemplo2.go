package main

import "fmt"

func main()  {
	animals := []string {
		"vaca",
		"cachorro",
		"falcao",
	}

	fmt.Println("o animal que voa é o:", animals[len(animals)])
}
