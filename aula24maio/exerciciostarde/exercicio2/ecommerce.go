package main

import "fmt"

/*
Exercício 2 - E-commerce (Parte II)
Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar produtos aos usuários.
Para fazer isso, eles exigem que usuários e produtos tenham o mesmo endereço de memória no main do programa e nas funções.

Estruturas necessárias:
- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
- Produto: Nome, preço, quantidade.

Algumas funções necessárias:
- Novo produto: recebe nome e preço, e retorna um produto.
- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona o produto ao usuário.
- Deletar produtos: recebe um usuário, apaga os produtos do usuário.

*/

type user struct {
	name string
	lastName string
	email    string
	products listProducts
}

type listUsers []user

type product struct {
	name string
	price float64
	quantity int
}

type listProducts []*product

//Novo produto: recebe nome e preço, e retorna um produto
func newProduct(name string, price float64) *product {
	return &product{
		name:  name,
		price: price,
	}
}

//Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona o produto ao usuário.
func addProduct(u *user, prod *product, quantity int) {
	prod.quantity = quantity
	u.products = append(u.products, prod)
}

//Deletar produtos: recebe um usuário, apaga os produtos do usuário.
func deleteProducts(u *user){
	u.products = u.products[:0]
}

func main()  {

	user1 := user{name: "ana", lastName: "silva", email: "ana.gmail.com", products: listProducts{}}
	prod1 := newProduct("banana", 5.00)
	prod2 := newProduct("melancia", 10.00)

	addProduct(&user1, prod1, 5)
	addProduct(&user1, prod2, 2)
	for _, prod := range user1.products {
		fmt.Println("Product: ",*prod)
	}

	deleteProducts(&user1)
	fmt.Println("After delete: ")

	prod3 := newProduct("abacaxi", 15.00)
	addProduct(&user1, prod3, 5)

	for _, prod := range user1.products {
		fmt.Println("Product: ",*prod)
	}
}