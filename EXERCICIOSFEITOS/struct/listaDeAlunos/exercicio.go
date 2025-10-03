package main

import "fmt"

type Aluno struct {
	Nome     string
	Nota     float64
	Aprovado bool
}

func main() {
	// Cria uma slice de structs do tipo Aluno e já inicializa com 3 alunos.
	alunos := []Aluno{
		{Nome: "Gabriela", Nota: 10, Aprovado: true},
		{Nome: "Eliana", Nota: 5, Aprovado: false},
		{Nome: "Carol", Nota: 8, Aprovado: true},
	}

	for _, aluno := range alunos {
		fmt.Printf("Nome: %s, Nota: %.1f, Aprovado: %t\n", aluno.Nome, aluno.Nota, aluno.Aprovado)
	}
}

/*Lista de structs
Crie uma slice de Aluno e adicione 3 alunos.
Imprima todos usando um for.*/
