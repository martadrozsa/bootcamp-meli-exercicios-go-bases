package main

import "fmt"

/*
Exercício 1 - Imprimindo o nome na tela
1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
2. Imprima no terminal o valor de cada variável.
*/

var (
	nome  string
	idade int
)

func main() {

	nome = "Marta"
	idade = 32
	fmt.Printf("Nome: %s\nIdade: %d", nome, idade)
}
