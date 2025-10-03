package main

import (
	"fmt"
)

type Funcionario struct {
	Nome  string
	Ativo bool
}

func criarFuncionario(nome string, ativo bool) Funcionario {
	return Funcionario{Nome: nome, Ativo: ativo}
}

func funcionarioAtivo(listaFuncionario []Funcionario) {
	for _, f := range listaFuncionario {
		if f.Ativo {
			fmt.Println(f.Nome)
		}
	}
}

func main() {

	listaFuncionario := []Funcionario{
		criarFuncionario("Eliana", true),
		criarFuncionario("Gabriela", false),
		criarFuncionario("Carlos", true),
		criarFuncionario("Mariana", false),
	}

	funcionarioAtivo(listaFuncionario)
}

/*Lista de funcionários ativos

Crie uma struct Funcionario com Nome e Ativo.
Crie uma slice de funcionários (alguns ativos e outros não).
Escreva uma função que receba essa slice e imprima somente os funcionários ativos.*/
