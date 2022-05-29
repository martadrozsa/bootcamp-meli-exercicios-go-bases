package main

import "fmt"

// Aumentar recebe um ponteiro de tipo inteiro
func Aumentar(v *int)  {
	//desreferenciamos a variável v para obter seu valor e incrementá-lo em 1
	*v++
}

func main()  {
	var v int = 19
	fmt.Println("O valor de v é:", v)
	// a função Aumentar recebe um ponteiro
	// utilizamos o operador de endereço & para passar o endereço de memória de v

	Aumentar(&v)
	fmt.Println("O valor de v agora é:", v)
}