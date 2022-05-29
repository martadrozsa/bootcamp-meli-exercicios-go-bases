package main

import "fmt"

/*
Atualizando valores
podemos atualizar o valor de um elemento específico consultando seu nome da chave
*/

func main() {
	var students = map[string]int{"João": 20, "Pedro": 26}
	fmt.Println(students)

	students["João"] = 22
	fmt.Println(students)
}