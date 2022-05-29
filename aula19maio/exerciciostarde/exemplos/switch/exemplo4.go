package main

import "fmt"

/*
Switch com declaração curta
*/

func main()  {

	switch dia := "terça-feira"; dia {
	case "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira":
		fmt.Printf("%s é um dia de semana\n", dia)
	default:
		fmt.Printf("%s é um dia de final de semana\n", dia)
	}
}
