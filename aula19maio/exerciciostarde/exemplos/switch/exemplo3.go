package main

import "fmt"

/*
Switch com múltiplos cases
*/

func main()  {

	dia := "domingo"

	switch dia {
	case "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira":
		fmt.Printf("%s é um dia de semana\n", dia)
	default:
		fmt.Printf("%s é um dia de final de semana\n", dia)
	}
}
