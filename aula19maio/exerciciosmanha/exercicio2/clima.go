package main

import "fmt"

/*
Exercício 2 - Clima
Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e pressão atmosférica de diferentes lugares.
1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
2. Imprima os valores no console.
3. Quais tipos de dados serão atribuídos a essas variáveis?
*/

var (
	temperatura        float64
	umidade            float64
	pressaoAtmosferica float64
)

func main() {
	temperatura = 17
	umidade = 63.4
	pressaoAtmosferica = 1009

	fmt.Printf("Temperatura: %.1f°\nUmidade: %.1f%%\nPressão atmosférica: %.1f hPa", temperatura, umidade, pressaoAtmosferica)
}
