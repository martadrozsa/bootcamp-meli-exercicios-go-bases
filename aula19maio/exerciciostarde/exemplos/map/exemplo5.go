package main

import "fmt"

/*
Deletando elementos
*/

func main() {
	var students = make(map[string]int)
	students["João"] = 20
	fmt.Println(students)

	delete(students, "João")
	fmt.Println(students)
}
