package main

import (
	"fmt"
	"time"
)

func processarComGoRoutine(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
}

func main()  {
	for i := 0; i < 10; i++ {
		go processarComGoRoutine(i)
	}

	//time.Sleep(5000 * time.Millisecond)
	//fmt.Println("Programa encerrado")
}