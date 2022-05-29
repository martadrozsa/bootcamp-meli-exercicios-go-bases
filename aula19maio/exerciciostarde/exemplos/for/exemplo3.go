package main

import "fmt"

/*
Loop infinito
são úteis dentro de rotinas Go, quando possuem um worker ou processo que deve continuar para sempre
*/

func main()  {
	sum := 0

	for  {
		sum++ // repetir para sempre
		fmt.Println(sum)
	}
}
//nunca termina, o loop é infinito