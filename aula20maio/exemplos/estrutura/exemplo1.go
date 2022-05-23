package main

import "fmt"

//definimos uma estrutura da seguinte forma: determinamos seus campos seguidos por um espaço e o tipo de dados

type Pessoa struct {
	Nome      string
	Genero    string
	Idade     int
	Profissao string
	Peso      float64
	Gostos    Preferencias
}

type Preferencias struct {
	Comidas  string
	Filmes   string
	Series   string
	Animes   string
	Esportes string
}

func main() {
	// instanciar uma estrutura
	p1 := Pessoa{"Maria", "F", 37, "Engenheira", 65.5, Preferencias{"Sorvete", "LoTR", "GoT", "Bleach", "Corrida"}}
	fmt.Println(p1)

	// instanciar uma estrutura
	p2 := Pessoa{
		Nome:      "Pedro",
		Idade:     25,
		Profissao: "Arquiteto",
		Peso:      77,
		Gostos: Preferencias{
			Comidas:  "Pizza",
			Filmes:   "Clube da luta",
			Series:   "The Office",
			Animes:   "Narutu",
			Esportes: "Natação",
		},
	}
	fmt.Println(p2)

	// acessar um campo
	fmt.Println(p2.Nome)

	// atribuir ou modificar um valor a um campo
	p2.Peso = 70
	fmt.Println(p2)

	p2.Gostos.Esportes = "Futebol"
	fmt.Println(p2)

	// definir uma estrutura vazia e depois atribuir valores
	var p3 Pessoa
	p3.Nome = "Ana"
	p3.Idade = 22
	fmt.Println(p3)
}
