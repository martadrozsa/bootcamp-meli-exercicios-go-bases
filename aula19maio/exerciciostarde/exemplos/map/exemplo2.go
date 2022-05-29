package main

import "fmt"

/*
Acessando elementos
Chamamos o nome do mesmo seguido pelo nome da chave que queremos acessar

uma chave funciona como um índice, apontando para o valor associado a essa chave
*/

func main() {
	var student = map[string]int{"João": 20, "Pedro": 26}
	fmt.Println(student["João"])
}