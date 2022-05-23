package main

import "fmt"

/*
Exercício 1 - Letras de uma palavra
A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma das letras separadamente para soletrá-la.
Para isso terão que:
1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de letras que ela contém.
2. Em seguida, imprimi cada uma das letras.
*/

func main() {

	palavra := "trabalho"
	fmt.Println("Quantidade de letras: ", len(palavra))

	for i := 0; i < len(palavra); i++ {
		fmt.Printf("%c\n", palavra[i])
	}

}
