package main

import "fmt"

/*
Exercício 2 - Empréstimo

Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido.
Apenas concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano de atividade.
Dentro dos empréstimos que concede, não cobra juros para quem tem um salário superior a US$ 100.000.

É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
*/

var idade int
var isEmpregado bool
var tempoDeAtividade int
var valorSalario float64

func main() {

	idade = 23
	isEmpregado = true
	tempoDeAtividade = 2
	valorSalario = 150000

	if idade > 22 && isEmpregado == true && tempoDeAtividade > 1 {

		if valorSalario > 100000 {
			fmt.Println("Empréstimo concedido sem juros")
		} else {
			fmt.Println("Empréstimo concedido")
		}

	} else {
		fmt.Println("Empréstimo não pode ser concedido")

	}
}
