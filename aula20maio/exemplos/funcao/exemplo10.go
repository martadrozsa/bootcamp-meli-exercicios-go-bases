package main

import "fmt"

/*
Retorno de funções
podemos implementar um função que retorne outra função, para isso devemos indicar os parâmetros e os tipo de dados que essa função retorna.

retornará outra função que recebe 2 parâmetros e retorna um valor de float
func minhaFunc(valor string) func(valor1, valor2 float64) float64 {}

*/

func operacaoSoma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func operacaoSubtra(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func operacaoMulti(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func operacaoDivi(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

//função que se encarrega de lidar com as funções que realizarão as operações
func operacoesAritmeticas(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case "Soma":
		return operacaoSoma
	case "Subtração":
		return operacaoSubtra
	case "Multiplicação":
		return operacaoMulti
	case "Divisão":
		return operacaoDivi
	}
	return nil
}

func main() {
	operacao := operacoesAritmeticas("Soma")
	resultado := operacao(10, 5)
	fmt.Println(resultado)
}
