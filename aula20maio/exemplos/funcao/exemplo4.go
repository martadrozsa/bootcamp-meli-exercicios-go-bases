package main

import "fmt"

const (
	Soma          = "+"
	Subtracao     = "-"
	Multiplicacao = "*"
	Divisao       = "/"
)

func operacaoAritmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Soma:
		return valor1 + valor2
	case Subtracao:
		return valor1 - valor2
	case Multiplicacao:
		return valor1 * valor2
	case Divisao:
		return valor1 / valor2
	}
	return 0
}

func main() {
	fmt.Println(operacaoAritmetica(6, 2, Soma))
	fmt.Println(operacaoAritmetica(6, 2, Subtracao))
	fmt.Println(operacaoAritmetica(6, 2, Multiplicacao))
	fmt.Println(operacaoAritmetica(6, 2, Divisao))
}
