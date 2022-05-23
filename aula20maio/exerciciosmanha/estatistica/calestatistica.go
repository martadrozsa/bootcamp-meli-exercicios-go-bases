package main

import (
	"errors"
	"fmt"
	"math"
)

/*
Exercício 4 - Cálculo de estatísticas

Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas dos alunos de um curso,
sendo necessário calcular os valores mínimo, máximo e médio de suas notas.

Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo, máximo ou média)
e que retorna outra função (e uma mensagem caso o cálculo não esteja definido)

que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi indicado na função anterior.
Exemplo:

const (
minimum = "minimum"
average = "average"
maximum = "maximum"
)
...
minhaFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...
minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calMinimum(notas ...float64) (float64, error) {
	if len(notas) == 0 {
		return 0.0, errors.New("inserir um conjunto de notas")
	}

	var menor float64 = math.MaxFloat64
	for i := 0; i < len(notas); i++ {
		if notas[i] < menor {
			menor = notas[i]
		}
	}
	return menor, nil
}

func calMaximum(notas ...float64) (float64, error) {
	if len(notas) == 0 {
		return 0.0, errors.New("inserir um conjunto de notas")
	}

	var maior float64 = math.MinInt64
	for i := 0; i < len(notas); i++ {
		if notas[i] > maior {
			maior = notas[i]
		}
	}
	return maior, nil
}

func calAverage(notas ...float64) (float64, error) {

	if len(notas) == 0 {
		return 0.0, errors.New("inserir um conjunto de notas")
	}
	soma := 0.0
	for i := 0; i < len(notas); i++ {
		soma = soma + notas[i]
	}
	return soma / float64(len(notas)), nil
}

// função que indica qual o tipo de cálculo que deve ser realizado
// e que retorna outra função (e uma mensagem caso o cálculo não esteja definido) que pode ser passado uma
// quantidade N de inteiros e retorne o cálculo que foi indicado na função anterior.

func tipoDeCalculo(tipoOperacao string) (
	func(notas ...float64) (float64, error),
	error) {

	switch tipoOperacao {
	case "minimum":
		return calMinimum, nil
	case "maximum":
		return calMaximum, nil
	case "average":
		return calAverage, nil
	default:
		return nil, errors.New("A operação não existe")
	}
}

func main() {

	operacaoMin, _ := tipoDeCalculo(minimum)
	resultMin, _ := operacaoMin(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println("Nota mínima:", resultMin)

	operacaoMax, _ := tipoDeCalculo(maximum)
	resultMax, _ := operacaoMax(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("Nota máxima:", resultMax)

	operacaoAve, _ := tipoDeCalculo(minimum)
	resultAve, _ := operacaoAve(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("Média:", resultAve)

}
