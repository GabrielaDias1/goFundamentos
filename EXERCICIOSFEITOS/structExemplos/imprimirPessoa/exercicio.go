package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// Função que recebe uma Pessoa e imprime no formato solicitado
func imprimirPessoa(p Pessoa) {
	fmt.Printf("Nome: %s, Idade: %d anos\n", p.Nome, p.Idade)
}

func main() {
	pessoa1 := Pessoa{
		Nome:  "Gabriela",
		Idade: 22,
	}

	// Chamando a função
	imprimirPessoa(pessoa1)
}

/*Função para imprimir

Crie uma struct Pessoa com Nome e Idade.
Depois, crie uma função que receba uma Pessoa e imprima no formato:
"Nome: X, Idade: Y anos".*/
