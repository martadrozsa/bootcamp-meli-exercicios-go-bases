package main

import "fmt"

/*
Adicionando elementos
utilizamos uma nova chave de índice e inserimos um valor
*/

func main() {
	var students = map[string]int{"João": 20, "Pedro": 26}
	fmt.Println(students)

	students["Brenda"] = 19
	students["Marcos"] = 22

	fmt.Println(students)
}