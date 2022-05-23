package main

import "fmt"

/*
podemos utilizar notação de reticências, permitindo que nossas funções recebam um número dinãmico de parâmetros
*/

func funcSoma(valores ...float64) float64 {
	var resultado float64
	for _, valor := range valores {
		resultado += valor
	}
	return resultado
}

func main() {
	s := funcSoma(10, 10, 10, 10, 10)
	fmt.Println(s)
}
