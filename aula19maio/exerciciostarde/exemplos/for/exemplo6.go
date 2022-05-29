package main

import "fmt"

/*
Pular para a próxima iteração
Passar para a próxima iteração antes que todo o código do loop termine de ser executado
*/

func main()  {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}
}
