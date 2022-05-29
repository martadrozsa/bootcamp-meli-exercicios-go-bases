package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Exercício 1 — Guardar arquivo
Uma empresa que vende produtos de limpeza precisa:
1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação de produtos comprados, separados por ponto e vírgula (csv).
2. Deve possuir o ID do produto, preço e a quantidade.
3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

Exercício 2 — Ler arquivo
A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
Seja impresso na tela os valores tabelados, com título (à esquerda para o ID e à direita para o Preço e Quantidade),
o preço, a quantidade e abaixo do preço o total deve ser exibido (somando preço por quantidade)

Exemplo:
ID 				Preco 			Quantidade
111223 			30012.00 			1
444321 			1000000.00 			4
434321 			50.50 				1
				4030062.50
*/

type product struct {
	id int
	price float64
	quantity float64
}

func (p product) print() {
	fmt.Print(p.id, "\t", p.price, "\t", p.quantity,"\n")
}

type fileManager interface {
	addData(data product) error
	readData() ([]product, error)
}

type csvManager struct {
	filePath string
}

func (c csvManager) addData(data product) error {
	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(c.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	productValues := fmt.Sprint(data.id, ";", data.price, ";", data.quantity, "\n")

	content := []byte(productValues)
	_, err = file.Write(content)
	if err != nil {
		file.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	return err
}

func (c csvManager) readData() ([]product, error) {

	//Lê o conteúdo do arquivo. Retorna no formato de um array de ‘bytes’
	data, err := os.ReadFile(c.filePath)
	if err != nil {
		return nil, err
	}

	//converte o array de bytes em uma string
	content := string(data)

	//fmt.Println("teste", content)

	//separa todas as linhas do arquivo em um slice para poder acessar linha por linha
	lines := strings.Split(content, "\n")

	//cria um slice de produtos para conter os produtos parseados das linhas do arquivo — cada linha tem os dados de um produto
	products := make([]product, 0)

	//para cada linha
	for _, line := range lines {

		//separa os elementos da linha usando o separador ;
		elements := strings.Split(line, ";")
		if len(elements) < 3 {
			//fmt.Println("Skiping linvalid line")
			continue
		}
		//fmt.Println("elements: ", elements)

		//atribui cada elemento a uma variável, convertendo de ‘string’ para outro tipo (quando necessário)
		id, _ := strconv.Atoi(elements[0])
		price, err := strconv.ParseFloat(elements[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		quantity, _ :=  strconv.ParseFloat(elements[2], 64)

		//cria uma instância de produto com os dados estraídos da linha
		p := product{id, price, quantity}

		products = append(products, p)
	}

	return products, nil
}

func createFileManager(filePath string) fileManager  {
	return csvManager{filePath}
}


func main()  {
	filePath := "./products.csv"

	var productsFileManager fileManager
	productsFileManager = createFileManager(filePath)

	product1 := product{1115, 150.00, 4}
	product2 := product{3256, 250.00, 2}
	product3 := product{2426, 130.00, 1}

	err := productsFileManager.addData(product1)
	if err != nil {
		return
	}
	err = productsFileManager.addData(product2)
	if err != nil {
		return
	}
	err = productsFileManager.addData(product3)
	if err != nil {
		return
	}


	datas, err := productsFileManager.readData()

	fmt.Print("\nID\t")
	fmt.Print("Price\t")
	fmt.Print("Quantity\t\n")

	var sum float64 = 0.0
	for _, p := range datas {
		sum += p.price * p.quantity
		p.print()
	}
	fmt.Print("\t", sum)
}