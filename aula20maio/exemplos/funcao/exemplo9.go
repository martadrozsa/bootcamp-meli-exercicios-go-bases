package main

import "fmt"

/*
Também podemos retornar valores nomeados
devemos definir na função o tipo de dado a ser retornado e também o nome da variável

func operacao(valor1, valor2 float64) (soma, subra, mult, divis float64) {}
*/

func operacao(valor1, valor2 float64) (soma, subra, mult, divis float64) {
	soma = valor1 + valor2
	subra = valor1 - valor2
	mult = valor1 * valor2
	divis = valor1 / valor2

	if valor2 != 0 {
		divis = valor1 / valor2
	}
	return
}

func main() {
	fmt.Println(operacao(25, 5))
}
