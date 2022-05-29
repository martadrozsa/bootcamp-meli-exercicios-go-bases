package main

import "fmt"

func main()  {
	name, age := "Kim", 22
	res := fmt.Sprint(name, " tem ", age, " anos de idade.\n")
	fmt.Print(res)

	res = fmt.Sprintf( "%s tem %d anos de idade.\n", name, age)
	fmt.Print(res)
}