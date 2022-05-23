package main

import "fmt"

/*
Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar produtos e devolver o valor do preço total.
As empresas têm 3 tipos de produtos:
- Pequeno
- Médio
- Grande.

Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

Lista de custos adicionais:
- Pequeno: O custo do produto (sem custo adicional)
- Médio: O custo do produto + 3% pela disponibilidade no estoque
- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo adicional pelo envio de $2500.

Requisitos:
- Criar uma estrutura “loja” que guarde uma lista de produtos.
- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
- Criar uma interface “Produto” que possua o método “CalcularCusto”
- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.

- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome e preço, e devolva um Produto.
- Será necessário uma função “novaLoja” que retorne um Ecommerce.

- Interface Produto:
- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o custo adicional segundo o tipo de produto.

- Interface Ecommerce:
- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com base no custo total dos produtos + o adicional citado anteriormente (caso a categoria tenha)
- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto e adicioná-lo a lista da loja
*/

type tipoProduto int

const (
	Pequeno tipoProduto = iota
	Medio
	Grande
)

// exemplo: produtoTec, produtoCamaMesaBanho, produtoVestuario
type produto struct {
	tipo  tipoProduto
	nome  string
	preco float64
}

func (p produto) calcularCusto() float64 {
	switch p.tipo {
	case Pequeno:
		return p.preco
	case Medio:
		return p.preco + p.preco*0.03
	case Grande:
		return p.preco + p.preco*0.06 + 2500
	default:
		return 0
	}
}

type Produto interface {
	calcularCusto() float64
}

type loja struct {
	listaProdutos []Produto
}

func (l loja) total() float64 {
	var total = 0.0
	for i := 0; i < len(l.listaProdutos); i++ {
		currentProduto := l.listaProdutos[i]
		total += currentProduto.calcularCusto()
	}
	return total
}

func (l *loja) adicionar(prod Produto) {
	l.listaProdutos = append(l.listaProdutos, prod)
}

type Ecommerce interface {
	total() float64
	adicionar(prod Produto)
}

//devolve um Produto(interface implementada pelo type produto)
func novoProduto(tipo tipoProduto, nome string, preco float64) Produto {
	return produto{
		tipo:  tipo,
		nome:  nome,
		preco: preco,
	}
}

func novaLoja() Ecommerce {
	return &loja{listaProdutos: make([]Produto, 0)}
}

func main() {

	loja := novaLoja()

	prodA := novoProduto(Pequeno, "Ceular", 2000)
	loja.adicionar(prodA)

	prodB := novoProduto(Medio, "Cadeira", 1000)
	loja.adicionar(prodB)

	prodC := novoProduto(Grande, "Sofá", 3000)
	loja.adicionar(prodC)

	totalCompra := loja.total()

	fmt.Println(totalCompra)

}
