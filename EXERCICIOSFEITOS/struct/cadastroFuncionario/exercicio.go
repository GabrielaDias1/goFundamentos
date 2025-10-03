package main

import "fmt"

type Funcionario struct {
	Nome    string
	Salario float64
	Ativo   bool
}

func main() {
	funcionarios1 := Funcionario{
		Nome:    "Gabriela",
		Salario: 1.00,
		Ativo:   true,
	}
	fmt.Printf("Funcionário %s tem salário %.2f e está ativo = %t\n", funcionarios1.Nome, funcionarios1.Salario, funcionarios1.Ativo)
}

/*Cadastro de funcionário

Crie uma struct Funcionario com Nome (string), Salario (float64) e Ativo (bool).
Crie um funcionário e mostre a mensagem:
"Funcionário X tem salário Y e está ativo = Z".*/
