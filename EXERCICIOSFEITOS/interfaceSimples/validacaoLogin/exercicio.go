package main

import "fmt"

type Login interface {
	Validar(usuario, senha string) bool
}

type Sistema struct{}

func (s Sistema) Validar(usuario, senha string) bool {
	return usuario == "admin" && senha == "1234"
}

func genericaInterface(l Login, usuario, senha string) {
	validado := l.Validar(usuario, senha)
	if validado {
		fmt.Println("Login com sucesso!")
	} else {
		fmt.Println("Usuário ou senha incorretos.")
	}
}

func main() {
	s := Sistema{}
	genericaInterface(s, "admin", "1234")
	genericaInterface(s, "admin", "errado")
}

/*Descrição:
Crie uma interface Login com o método:

Validar(usuario, senha string) bool

Implemente em uma struct Sistema.

Considere usuário "admin" e senha "1234".

Teste no main com usuário e senha diferentes.*/
