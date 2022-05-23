package main

import (
	"errors"
	"fmt"
)

/*
Exercício 5 - Cálculo da quantidade de alimento
Um abrigo de animais precisa descobrir quanta comida comprar para os animais de estimação.
No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão é que haja muito mais animais para abrigar.
1. Cães precisam de 10 kg de alimento
2. Gatos 5 kg
3. Hamster 250 gramas.
4. Tarântula 150 gramas.

Precisamos:
1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto com o animal especificado e que retorne uma função com uma mensagem
(caso não exista o animal)
2. Uma função para cada animal que calcule a quantidade de alimento com base na quantidade necessária do animal digitado.
Exemplo:

const (
dog = "dog"
cat = "cat"
)

...
animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)

...
var amount float64
amount+= animaldog(5)
amount+= animalCat(8)
*/

const (
	tarantulas = "tarantulas"
	hamsters   = "hamsters"
	cachorro   = "cachorro"
	gato       = "gato"
)

func calCachorro(quantidadeAninal int) int {
	return quantidadeAninal * 10
}

func calTarantulas(quantidadeAninal int) int {
	return quantidadeAninal * 150
}

func calHamsters(quantidadeAninal int) int {
	return quantidadeAninal * 250
}

func calGato(quantidadeAninal int) int {
	return quantidadeAninal * 5
}

func animal(animal string) (
	func(quantidadeAninal int) int, error) {
	switch animal {
	case tarantulas:
		return calTarantulas, nil
	case cachorro:
		return calCachorro, nil
	case hamsters:
		return calHamsters, nil
	case gato:
		return calGato, nil
	default:
		return nil, errors.New("insira um animal válido")
	}
}

func main() {
	animalTaran, _ := animal(tarantulas)
	fmt.Printf("Quantidade de alimento para a tarântula em gramas: %d gramas\n", animalTaran(7))

	animalCachorro, _ := animal(cachorro)
	fmt.Printf("Quantidade de alimento para o cachorro em kilo: %d kilos\n", animalCachorro(8))

	animalHamsters, _ := animal(hamsters)
	fmt.Printf("Quantidade de alimento para o hamsters em gramas: %d gramas\n", animalHamsters(12))

	animalGato, _ := animal(gato)
	fmt.Printf("Quantidade de alimento para o gato em kilo: %d kilos\n", animalGato(6))

}
