package main

import "fmt"

/*
O conceito de herança não existe em GO, mas temos uma composição que usa a estrutura pai como um campo em nossas
estruturas filhas, isso é conhecido como embedding structs
*/

type Veiculo struct {
	km   float64
	hora float64
}

type Automovel struct {
	v Veiculo
}

type Moto struct {
	v Veiculo
}

func (v Veiculo) printDetalhe() {
	fmt.Printf("km:\t%f\nhora:\t%f\n", v.km, v.hora)
}

func (auto *Automovel) Correr(minutos int) {
	auto.v.hora = float64(minutos) / 60
	auto.v.km = auto.v.hora * 100
}

func (auto *Automovel) Detalhe() {
	fmt.Println("\nV:\tAutomovel")
	auto.v.printDetalhe()
}

func (moto *Moto) Correr(minutos int) {
	moto.v.hora = float64(minutos) / 60
	moto.v.km = moto.v.hora * 80
}

func (moto *Moto) Detalhe() {
	fmt.Println("\nV:\tAutomovel")
	moto.v.printDetalhe()
}

func main() {
	automovel := Automovel{}
	automovel.Correr(360)
	automovel.Detalhe()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalhe()
}
