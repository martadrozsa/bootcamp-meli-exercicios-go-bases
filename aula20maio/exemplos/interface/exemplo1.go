package main

import (
	"fmt"
	"math"
)

/*
uma interface é uma forma de definir métodos que devem ser usados, mas sem defini-los
são usadas para fornecer modularidade à linguagem
*/

type geometria interface {
	area() float64
	perimetro() float64
}

type triangulo struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (t triangulo) area() float64 {
	return t.largura * t.altura
}

func (t triangulo) perimetro() float64 {
	return 2*t.largura + 2*t.altura
}

func printDetalhe(g geometria) {
	fmt.Println(g)
	fmt.Printf("Área: %.2f\n", g.area())
	fmt.Printf("Perímetro: %.2f\n", g.perimetro())
}

func main() {
	c := circulo{raio: 5}
	printDetalhe(c)

	t := triangulo{altura: 2, largura: 3}
	printDetalhe(t)
}
