package main

/*
Exercício 3 - Calcular Preço (Part II)
Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total dos Produtos, Serviços e Manutenção.
Devido à forte demanda e para otimizar a velocidade, eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

Precisamos de 3 estruturas:
- Produtos: nome, preço, quantidade.
- Serviços: nome, preço, minutos trabalhados.
- Manutenção: nome, preço.
Precisamos de 3 funções:
- Somar Produtos: recebe um array de produto e devolve o preço total (preço * quantidade).
- Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média hora trabalhada,
se ele não vier trabalhar por 30 minutos, ele será cobrado como se tivesse trabalhado meia hora).
- Somar Manutenção: recebe um array de manutenção e devolve o preço total.

Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na tela (somando o total dos 3).
*/

type product struct {
	name string
	price float64
	quantity float64
}

type service struct {
	name string
	price float64
	minutesWorked float64
}

type keeping struct {
	name string
	price float64
}

type listProducts []product

type listServices []service

type listKeeping []keeping


// Somar Produtos: recebe um array de produto e devolve o preço total (preço * quantidade)
func sumProducts(products listProducts) float64 {
	return 0
}

// Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média hora trabalhada,
// se ele não vier trabalhar por 30 minutos, ele será cobrado como se tivesse trabalhado meia hora).
func sumServices(services listServices) float64  {
	return 0
}

// Somar Manutenção: recebe um array de manutenção e devolve o preço total.
func sumKeeping(keeping listKeeping) float64  {
	return 0
}