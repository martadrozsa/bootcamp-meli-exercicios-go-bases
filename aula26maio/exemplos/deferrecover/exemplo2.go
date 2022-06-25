package main

import "fmt"

func isEven(num int) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	if (num % 2) != 0 {
		panic("não é par")
	}

	fmt.Println(num, "é um número par")
}

func main()  {
	num := 3
	isEven(num)

	fmt.Println("Execução finalizada com sucesso")
}