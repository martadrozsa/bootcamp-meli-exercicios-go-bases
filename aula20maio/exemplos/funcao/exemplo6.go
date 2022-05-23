package main

import "fmt"

// constantes com as operações a serem realizadas
const (
	Sum   = "+"
	Subt  = "-"
	Multi = "*"
	Divis = "/"
)

// funcao para realizar a operação Soma
func opeSum(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func opeSubt(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func opeMulti(valor1, valor2 float64) float64 {
	return valor1 * valor2
}
func opeDivis(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

// função que retornará as operações
func retornoOperacao(valores []float64, operacao func(valor1, valor2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacao(resultado, valor)
		}
	}
	return resultado
}

// função que ficará encarregada de receber a operação a ser realizada e os valores aos quais a operação será aplicada
// para cada operação será chamada uma função que recebe os valores e a função que vamos executar para aquele operador
func opeAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Sum:
		return retornoOperacao(valores, opeSum)
	case Subt:
		return retornoOperacao(valores, opeSubt)
	case Multi:
		return retornoOperacao(valores, opeMulti)
	case Divis:
		return retornoOperacao(valores, opeDivis)
	}
	return 0
}

func main() {
	fmt.Println(opeAritmetica(Sum, 10, 10, 10, 10, 10, 10, 10))
}
