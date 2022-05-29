package main

import (
	"fmt"
	"time"
)

// receber o canal como parâmetro na nossa função
func processarComGoRoutineComChanel(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")

	// valor de i será enviado para o canal
	c <- i
}

func main()  {

	c := make(chan int)
	for i := 0; i < 10; i++ {
		go processarComGoRoutineComChanel(i, c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Programa encerrado", <- c)
	}

}