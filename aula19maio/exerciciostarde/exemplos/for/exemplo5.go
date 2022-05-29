package main

import "fmt"

/*
Quebrar um laço de repetição
A palavra reservada break nos permite cortar com a execução do loop
*/

func main()  {
	sum := 0
	for  {
		sum++
		if sum >= 1000 {
			break
		}
	}
	fmt.Println(sum)
}
