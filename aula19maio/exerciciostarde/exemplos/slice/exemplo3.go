package main

import "fmt"

/*
Comprimento e capacidade
podem ser obtido usando as funções len(s) e cap(s)

Adcionando valores em um Slice
- append
*/


func main()  {

	var s []int
	s = append(s, 2, 3, 4)

	fmt.Println(s)
}


