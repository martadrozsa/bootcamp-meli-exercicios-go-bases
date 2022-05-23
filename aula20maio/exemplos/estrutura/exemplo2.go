package main

import (
	"encoding/json"
	"fmt"
)

/*
Rótulos e estruturas
Dentro de nossas estuturas podemos definir rótulos ou anotações que se referem a cada um dos campos que aparecem após a declaração do tipo de dado
type MinhaEstrutura struct {
	Campo1 string `meuRotulo: "valor"`
	Campo2 string `meuRotulo: "valor"`
}
*/

// Pessoa1 exemplo utilizando o formato JSON
type Pessoa1 struct {
	PrimeiroNome string `json:"primeiro_nome"`
	Sobrenome    string `json:"sobrenome"`
	Telefone     string `json:"telefone"`
	Endereco     string `json:"endereco"`
}

func main() {
	p := Pessoa1{"Ana", "Santos", "66666666", "Rua Bento, 123"}
	converteJson, erro := json.Marshal(p)
	fmt.Println(string(converteJson))
	fmt.Println(erro)
}
