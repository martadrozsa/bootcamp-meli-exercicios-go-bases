package main

import "fmt"

/*
Exercício 3 - Calcular salário
Uma empresa marítima precisa calcular o salário de seus funcionários com base no número de horas trabalhadas por mês e na categoria do funcionário.

Se a categoria for C, seu salário é de R$1.000 por hora
Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h mensais
Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h mensais

Calcule o salário dos funcionários conforme as informações abaixo:
- Funcionário de categoria C: 162h mensais
- Funcionário de categoria B: 176h mensais
- Funcionário de categoria A: 172h mensais
*/

func calCatA(horas, salario float64) float64 {
	if horas > 160 {
		return horas*salario + (0.50 * salario)
	} else {
		return salario * horas
	}
}

func calCatB(horas, salario float64) float64 {
	if horas > 160 {
		return horas*salario + (0.20 * salario)
	} else {
		return salario * horas
	}
}

func calCatC(horas, salario float64) float64 {
	return salario * horas
}

func main() {
	fmt.Printf("Salário para a categoria A: R$ %2.f\n", calCatA(172, 3000))
	fmt.Printf("Salário para a categoria B: R$ %2.f\n", calCatB(176, 1500))
	fmt.Printf("Salário para a categoria C: R$ %2.f", calCatC(122, 1000))
}
