package main

import (
	"fmt"
	"time"
)

/*
Switch com fallthrough
*/

func main()  {
	hoje := time.Now()
	var t int = hoje.Day()

	switch t {
	case 5, 10, 15:
		fmt.Println("Limpar a casa.")
	case 25, 26, 27:
		fmt.Println("Comprar comida.")
		fallthrough
	case 31:
		fmt.Println("Hoje tem festa")
	default:
		fmt.Println("Não há informações disponíveis para esse dia")



	}
}
