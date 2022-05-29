package main

import "fmt"

/*
Loop while -> permite que você execute um bloco de código até que a condição não seja mais verdadeira
feito apenas com um componente, a condição
GO não tem palavra-chave while para este loop
*/

func main()  {
	sum := 1

	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}