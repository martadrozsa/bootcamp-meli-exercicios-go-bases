package main

import "fmt"

// criando slice com make()
// a função make() gera um array com os valores em 0 e devolve uma slice que faz referência a esse array

func main()  {
	a := make([]int, 5) //len (a) = 5

	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4]) // obter valores de um intervalo dentro do slice
}


