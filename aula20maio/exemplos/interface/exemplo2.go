package main

import (
	"fmt"
	"math"
)

type geometricFigure interface {
	area() float64
	perimetro() float64
}

const (
	typeCircle   = "circle"
	typeTriangle = "triangle"
)

func newFigure(geoFigure string, values ...float64) geometricFigure {
	switch geoFigure {
	case typeCircle:
		return circle{raio: values[0]}
	case typeTriangle:
		return triangle{width: values[0], height: values[1]}
	}
	return nil
}

type circle struct {
	raio float64
}
type triangle struct {
	width  float64
	height float64
}

func (t triangle) area() float64 {
	return t.width * t.height
}

func (t triangle) perimetro() float64 {
	return 2*t.width + 2*t.height
}

func (c circle) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circle) perimetro() float64 {
	return 2 * math.Pi * c.raio
}
func main() {
	c := newFigure(typeCircle, 2)
	fmt.Printf("%2.f\n", c.area())
	fmt.Printf("%2.f\n", c.perimetro())

	t := newFigure(typeTriangle, 2, 3)
	fmt.Printf("%2.f\n", t.area())
	fmt.Printf("%2.f\n", t.perimetro())

}
