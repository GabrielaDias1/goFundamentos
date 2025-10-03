package main

import "fmt"

type Funcionario struct {
	Nome  string
	Cargo string
}

type Departamento struct {
	NomeDepto   string
	Funcionario []Funcionario
}

type Empresa struct {
	Nome         string
	Departamento []Departamento
}

func main() {
	empresaSamsung := Empresa{
		Nome: "Samsung",
		Departamento: []Departamento{
			{
				NomeDepto: "TI",
				Funcionario: []Funcionario{
					{Nome: "Gabriela", Cargo: "Estágio"},
					{Nome: "Elias", Cargo: "Estágio"},
				},
			},
			{
				NomeDepto: "RH",
				Funcionario: []Funcionario{
					{Nome: "Laura", Cargo: "Estágio"},
					{Nome: "Gabriel", Cargo: "Estágio"},
				},
			},
		},
	}

	
	fmt.Printf("Empresa: %s\n", empresaSamsung.Nome)

	
	for _, depto := range empresaSamsung.Departamento {
		fmt.Printf(" Departamento: %s\n", depto.NomeDepto)

		for _, f := range depto.Funcionario {
			fmt.Printf("   - %s (%s)\n", f.Nome, f.Cargo)
		}
	}
}

/*Empresa com departamentos e funcionários

Crie uma struct Funcionario com Nome e Cargo.

Crie uma struct Departamento com NomeDepto e slice de Funcionario.

Crie uma struct Empresa com Nome e slice de Departamento.

No main, crie uma empresa com 2 departamentos e pelo menos 2 funcionários em cada. Imprima a estrutura hierárquica no console.*/
