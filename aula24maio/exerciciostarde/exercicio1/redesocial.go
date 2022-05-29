package main

import "fmt"

/*
Exercício 1 — Rede social
Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções que acrescentem informações à estrutura.
Para otimizar e economizar memória, eles exigem que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e para as funções:
— A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e senha
E devem implementar as seguintes funcionalidades:
— mudar o nome: permite mudar o nome e sobrenome
— mudar a idade: permite mudar a idade
— mudar e-mail: permite mudar o e-mail
— mudar a senha: permite mudar a senha

*/

type usuario struct {
	id int
	nome string
	sobrenome string
	idade int
	email string
	senha string
}

type usuarios []*usuario

func (u *usuarios) addUsuario(usuario *usuario) {
	*u = append(*u, usuario)
}

func (u *usuarios) editUsuario(usuario *usuario, id int)  {
	for _, usuarioEdit := range *u {
		if usuarioEdit.id == id  {
			usuarioEdit.nome = usuario.nome
			usuarioEdit.sobrenome = usuario.sobrenome
			usuarioEdit.idade = usuario.idade
			usuarioEdit.email = usuario.email
			usuarioEdit.senha = usuario.senha
		} else {
			fmt.Println("Usuário não encontrado")
		}
	}
}

func (u usuarios) listUsuarios() {
	for _, usuario := range u {
		fmt.Println("Nome:", usuario.nome)
		fmt.Println("Sobrenome:", usuario.sobrenome)
		fmt.Println("Idade:", usuario.idade)
		fmt.Println("Email:", usuario.email)
		fmt.Println("Senha:", usuario.senha)
	}
}

func main() {
	u1 := usuarios{}
	user1 := usuario{id: 1, nome: "Lucas", sobrenome: "Lima", idade: 18, email: "lucas@gmail.com", senha: "lucas123"}

	u1.addUsuario(&user1)
	u1.listUsuarios()

	u2 := usuarios{}
	user2 := usuario{id: 2, nome: "Maria", sobrenome: "Alves", idade: 21, email: "maria@gmail.com", senha: "maria123"}

	u2.addUsuario(&user2)
	u2.listUsuarios()

	fmt.Println("\nUsuário 1 editado:")
	u1.editUsuario(&user1, 1)
	user1 = usuario{id: 1, nome: "Lucas", sobrenome: "Lima", idade: 19, email: "lucas.lima@gmail.com", senha: "lucas123456"}
	u1.listUsuarios()


	//fmt.Println(&u1)
	//fmt.Println(&user1)
	//fmt.Println(&u2)
	//fmt.Println(&user2)
}