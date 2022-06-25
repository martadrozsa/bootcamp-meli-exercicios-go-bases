package main

import "fmt"

type Dog struct {
	Name string
}

func (s *Dog) WoofWoof() {
	fmt.Println(s.Name, "faz woof woof")
}

func main()  {
	s := &Dog{"Samy"}
	s = nil
	s.WoofWoof()
}

/*
Ocorre um panic ao tentar acessar um endereço de memória inválido ou apontar para um ouvinte nulo
O panic ocorre tem tempo de execução e aborta a execução do programa com status 2
*/