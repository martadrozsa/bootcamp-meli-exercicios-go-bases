package main

import "fmt"

/*
Type assertion - declaração de tipo

a asserção de tipo fornece acesso ao tipo de dados exato que é abstraído por uma interface
*/

var i interface{} = "hello"

func main() {

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
