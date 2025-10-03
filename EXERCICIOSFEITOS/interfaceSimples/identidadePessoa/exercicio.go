package main

import "fmt"

type Identidade interface {
	Nome() string
	Idade() int
}

type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) Nome() string {
	return p.nome
}

func (p Pessoa) Idade() int {
	return p.idade
}

func criarIdentidade(i Identidade) {
	fmt.Println(i.Nome())
	fmt.Println(i.Idade())
}

func main() {
	gabriela := Pessoa{nome: "Gabriela", idade: 22}

	criarIdentidade(gabriela)
}

