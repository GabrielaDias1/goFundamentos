package main

import "fmt"

type Funcionario struct {
	Nome    string
	Salario float64
}

func criarFuncionario() []Funcionario {
	listaFuncionario := []Funcionario{
		{Nome: "Gabriela", Salario: 100.00},
		{Nome: "Giulia", Salario: 200.00},
		{Nome: "Fabiana", Salario: 300.00},
	}
	return listaFuncionario
}

func valorSalario(listaFuncionario []Funcionario) float64 {
	soma := 0.0
	for _, f := range listaFuncionario {
		soma += f.Salario
	}
	return soma
}

func main() {
	lista := criarFuncionario()               // pega a lista
	total := valorSalario(lista)              // calcula a soma
	fmt.Println("Total dos salários:", total) // imprime

}

/*Soma de salários

Crie uma struct Funcionario com Nome e Salario.
Crie uma slice com 3 funcionários.
Escreva uma função que calcule a soma dos salários e retorne o valor total.*/
