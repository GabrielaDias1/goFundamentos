package main

import "fmt"

type Aluno struct {
	Nome  string
	Idade int
}

type Turma struct {
	NomeTurma string
	Aluno     []Aluno
}

func main() {
	turma := Turma{
		NomeTurma: "Turma de matemática",
		Aluno: []Aluno{
			{Nome: "Gabriela", Idade: 22},
			{Nome: "Gabriel", Idade: 20},
			{Nome: "Ana", Idade: 21},
		},
	}

	fmt.Printf("%s tem os alunos:\n", turma.NomeTurma)
	for _, aluno := range turma.Aluno {
		fmt.Printf("- %s (%d anos)\n", aluno.Nome, aluno.Idade)
	}
}

/*Relacionamento entre structs

Crie uma struct Aluno com Nome e Idade.
Depois, crie uma struct Turma que tenha NomeTurma e uma slice de Aluno.
No main, crie uma turma com 3 alunos e imprima os dados no formato:
"Turma X tem os alunos: ...".*/
