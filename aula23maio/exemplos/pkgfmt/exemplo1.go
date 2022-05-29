package main

import "fmt"

func main()  {
	name, age := "Kim", 22
	fmt.Print(name, " tem ", age, " anos de idade.\n")

	fmt.Println(name, "tem", age, "anos de idade.")

	fmt.Printf("%10.2f\n", 12222.122123)
	fmt.Printf("%.2f\n", 12222.122123)

	fmt.Printf("%s tem %d anos de idade.\n", name, age)

	fmt.Printf("%10d", 12222)
	fmt.Printf("%10s", "aa")
}
