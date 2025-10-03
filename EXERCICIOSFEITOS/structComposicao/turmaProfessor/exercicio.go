package main

import "fmt"

type Professor struct {
	Nome       string
	Disciplina string
}

type Aluno struct {
	Nome string
}

type Turma struct {
	NomeTurma string
	Professor Professor
	Aluno     []Aluno
}
type Escola struct {
	Nome  string
	Turma []Turma
}

func main() {
	escola := Escola{
		Nome: "Escola Municipal de Artes",
		Turma: []Turma{
			{
				NomeTurma: "Turma A",
				Professor: Professor{
					Nome:       "Sonia",
					Disciplina: "Artes Iniciantes",
				},
				Aluno: []Aluno{
					{Nome: "Gabriela"},
					{Nome: "Isabela"},
				},
			},
			{
				NomeTurma: "Turma B",
				Professor: Professor{
					Nome:       "Artemis",
					Disciplina: "Artes Intermediário",
				},
				Aluno: []Aluno{
					{Nome: "Beatriz"},
					{Nome: "Carol"},
				},
			},
		},
	}

	fmt.Printf("Escola: %s\n", escola.Nome)

	for _, turma := range escola.Turma {
		fmt.Printf("\nTurma: %s\n", turma.NomeTurma)
		fmt.Printf(" Professor: %s (%s)\n", turma.Professor.Nome, turma.Professor.Disciplina)
		
		for _, aluno := range turma.Aluno {
			fmt.Printf("  - Aluno: %s\n", aluno.Nome)
		}
	}
}

/*Escola com turmas e professores

Crie uma struct Professor com Nome e Disciplina.

Crie uma struct Turma com NomeTurma e slice de Aluno.

Crie uma struct Escola com Nome e slice de Turma.

Crie uma escola com 2 turmas, cada uma com 2 alunos, e associe professores às turmas. Imprima:*/
