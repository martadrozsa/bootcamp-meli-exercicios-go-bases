package main

import "fmt"

/*
retorno de função
temos que indicar o tipo de dado que esperamos no final da função

criamos uma função que recebe dois parâmetros e retorna a soma deles
*/

func soma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func main() {
	s := soma(4, 5)
	fmt.Println(s)
}
