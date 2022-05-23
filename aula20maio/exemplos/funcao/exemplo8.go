package main

/*
funções de multi retorno
em GO o retorno multivalor é normalmente utilizado quando precisamos retornar um valor e um erro
e precisamos validar se ocorreu um erro ou não
*/

import (
	"errors"
	"fmt"
)

//implementamos nossa função de divisão e validamos se o divisor for zero, se for, retornará um erro, caso contrário fará a divisão
func divisao(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("O dividor não pode ser zero ")
	}
	return dividendo / divisor, nil
}

func main() {
	resultado, erro := divisao(10, 5)

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(resultado)
	}
}
