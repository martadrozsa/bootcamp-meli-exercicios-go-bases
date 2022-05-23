package main

import (
	"fmt"
	"math"
)

/*
Em GO os métodos podem ser definidos em tipos de dados
Um método é uma função com um argumento de receptor especial
O receptor aparece em sua própria lista de argumentos entre a palavra-chave func e o nome do método
func medoto(){}
*/

type Circulo struct {
	raio float64
}

//método que calcula a área
func (c Circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

//método que calcula o perímetro
func (c Circulo) perimetro() float64 {
	return math.Pi * c.raio
}

func main() {
	c1 := Circulo{raio: 5}
	fmt.Printf("Área do círculo: %.2f\n", c1.area())

	c2 := Circulo{raio: 5}
	fmt.Printf("Perímetro do círculo: %.2f", c2.perimetro())
}
