package main

import (
	"fmt"
	"math"
)

/*
Ponteiros
Quando criamos uma função e passamos uma variável como argumento, o que a função faz é copiar o valor da variável e trabalhar com esse valor.
A variável que passamos como argumento não é modificada

Ponteiros são muito úteis nos casos em que queremos passar uma variável como argumento para uma função para que seu valor seja modificado, é o que se conhece como
passar valores por referência a uma função.
*/
type Circle struct {
	raio float64
}

//método que calcula a área
func (c Circle) area() float64 {
	return math.Pi * c.raio * c.raio
}

//método que calcula o perímetro
func (c Circle) perimetro() float64 {
	return math.Pi * c.raio
}

func (c *Circle) setRaio(raio float64) {
	c.raio = raio
}
func main() {
	c := Circle{raio: 5}
	fmt.Printf("Área do círculo: %.2f\n", c.area())
	fmt.Printf("Perímetro do círculo: %.2f\n", c.perimetro())

	c.setRaio(10)
	fmt.Printf("Área do círculo: %.2f\n", c.area())
	fmt.Printf("Perímetro do círculo: %.2f", c.perimetro())
}
