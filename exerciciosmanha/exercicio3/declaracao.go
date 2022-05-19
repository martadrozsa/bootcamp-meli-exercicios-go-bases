package main

import "fmt"

/*
Exercício 3 - Declaração de variáveis
Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
Programação I para fornecer-lhes o feedback correspondente.
Um dos pontos do exame é declarar diferentes variáveis.
Ajude o professor com as seguintes questões:
1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
2. Corrigir as incorrectas.

var 1nome string -> incorreta
var sobrenome string -> correta
var int idade 	-> incorreta
1sobrenome := 6 -> incorreta
var licenca_para_dirigir = true -> incorreta
var estatura da pessoa int  -> incorreta
quantidadeDeFilhos := 2	-> correta
*/

var nome string
var idade int

var sobrenome = "Denisczwicz"

var licencaParaDirigir bool = true

var estaturaDaPessoa float64

func main() {
	quantidadeDeFilhos := 2
	fmt.Println(quantidadeDeFilhos)
}
