package main

import "fmt"

type Imprimivel interface {
	Imprimir() string
}

type Livro struct {
	titulo string
}

func (l Livro) Imprimir() string {
	return l.titulo
}

/*
Crie uma interface chamada Imprimivel com um método Imprimir() string.
Implemente essa interface em uma struct Livro que tem o campo titulo.
No método, retorne o título do livro.
*/
func main() {
	var i Imprimivel
	//Variável de interface → guarda qualquer tipo que respeite o contrato da interface.

	i = Livro{titulo: "Percy Jackson"}
	fmt.Println(i.Imprimir())
}
