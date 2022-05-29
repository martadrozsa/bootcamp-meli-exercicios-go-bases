package main

import "fmt"

/*
Comprimento de um map --> usando len()
*/

func main()  {
	var myMap = map[string]int{}
	fmt.Println(len(myMap)) //len() devolve 0 para um map não inicializado
}

