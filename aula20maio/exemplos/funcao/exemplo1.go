package main

import "fmt"

// uma função recebe um, muitos parâmetros ou nenhum
// pode retornar ou não um valor

func verificarVariavel(numero int) {
	if numero < 0 {
		fmt.Println("O número é negativo")
	} else if numero > 0 {
		fmt.Println("O número é positivo")
	} else {
		fmt.Println("O número é 0")
	}
}

func main() {
	a, b, c, d := 1, 0, 5, -3

	verificarVariavel(a)
	verificarVariavel(b)
	verificarVariavel(c)
	verificarVariavel(d)

}
