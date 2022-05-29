package main

import (
	"fmt"
	"time"
)

func processar(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
}

func main()  {
	for i := 0; i < 10; i++ {
		processar(i)
	}

	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Programa encerrado")
}