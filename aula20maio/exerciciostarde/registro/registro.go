package main

import "fmt"

/*
Exercício 1 - Registro de estudantes

Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada um deles, conforme o exemplo abaixo:
Nome: [Nome do aluno]
Sobrenome: [Sobrenome do aluno]
RG: [RG do aluno]
Data de admissão: [Data de admissão do aluno]

Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos alunos.
Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG, Data e que tenha um método de detalhamento

*/

type Aluno struct {
	Nome           string
	Sobrenome      string
	RG             int
	DataDeAdmissao string
}

func (a *Aluno) Detalhamento() {
	fmt.Printf("Nome: %s\nSobrenome: %s\nRG: %v\nData de Admissão: %s\n", a.Nome, a.Sobrenome, a.RG, a.DataDeAdmissao)
}

func main() {
	a := Aluno{
		Nome:           "Pedro",
		Sobrenome:      "Silva",
		RG:             12346,
		DataDeAdmissao: "20/05/2022",
	}
	a.Detalhamento()
}
