package main

import "fmt"

/*
Exercício 2 - Calcular média
Um colégio precisa calcular a média das notas (por aluno).
Precisamos criar uma função na qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
*/

func calcularMedia(notas ...int) int {
	soma := 0
	for i := 0; i < len(notas); i++ {
		soma = soma + notas[i]
	}

	return soma / int(len(notas))
}

func main() {
	fmt.Println("A média das notas é:", calcularMedia(6, 5, 7, 8, 9))
}
