package main

import "fmt"

/*
Podemos criar funções que retornam mais de um valor
func minhaFuncRetorno(valor1, valor2 float64) (float64, string, int, bool) {}
*/

func operacoes(valor1, valor2 float64) (float64, float64, float64, float64) {

	soma := valor1 + valor2
	subta := valor1 - valor2
	multip := valor1 * valor2
	divis := valor1 / valor2

	if valor2 != 0 {
		divis = valor1 / valor2
	}

	return soma, subta, multip, divis
}

func main() {
	soma, subtracao, multiplicacao, divisao := operacoes(6, 2)
	fmt.Println("Soma:\t\t", soma)
	fmt.Println("Subtração:\t", subtracao)
	fmt.Println("Multiplicação:\t", multiplicacao)
	fmt.Println("Divisão:\t", divisao)
}
