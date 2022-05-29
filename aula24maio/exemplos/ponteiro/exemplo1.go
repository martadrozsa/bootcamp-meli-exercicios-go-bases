package main

import "fmt"

/*
// * para criar variável do tipo pointeiro
// * para visualizar o valor da posição da memória
// & para atribuir um valor ao ponteiro
*/

func main()  {
	var p1 *int
	var p2 = new(int)
	p3 := &p1

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	var v int = 19
	fmt.Println("O endereço da variável v é:", &v)

	fmt.Println()
	var x int = 20
	var p *int // criamos um ponteiro p para referenciar o endereço na memória da variável x

	p = &x
	fmt.Printf("O ponteiro p refere-se ao endereço de memória: %v\n", p)
	fmt.Printf("Ao desreferenciar o ponteiro p, obtendo o valor: %d\n", *p)
}