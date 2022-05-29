package main

/*
arrays não podem ser redimencionados
*/

import "fmt"

var a [2]string

func main()  {
	//atribuindo valores
	a[0] = "Hello"
	a[1] = "World"

	//obtendo valores
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

