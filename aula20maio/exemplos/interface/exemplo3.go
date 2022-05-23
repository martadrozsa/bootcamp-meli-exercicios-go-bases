package main

import "fmt"

/*
Interfaces vazias
não possuem métodos declarados

são úteis para armazenar valores que são de um tipo de dado desconhecido, ou que podem varias de acordo com o fluxo do programa
*/

type ListaHeterogenea struct {
	Data []interface{}
}

func main() {
	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "olá")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)
}