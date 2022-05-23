package main

import "fmt"

/*
Exercício 1 - Impostos de salário

Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de depositar o salário,
para cumprir seu objetivo será necessário criar uma função que retorne o imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será descontado 17% do salário.
Um outro funcionário ganha um salário de $150.000, e nesse caso será descontado mais 10%.
*/

func calcularImposto(salario float64) float64 {
	if salario < 50000 {
		return 0.0
	} else if salario <= 150000 {
		return 0.27 * salario
	}
	return salario * 0.17
}

func main() {
	fmt.Printf("Valor do imposto para o salário de até RS50.000: R$ %.2f\n", calcularImposto(50000))
	fmt.Printf("Valor do imposto para o salário de até RS150.000: R$ %.2f\n", calcularImposto(150000))
	fmt.Printf("Valor do imposto para o salário acima de RS150.000: R$ %.2f", calcularImposto(160000))
}
