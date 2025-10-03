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
		NomeTurma: "Turma de Matemática",
		Aluno: []Aluno{
			{Nome: "Gabriela", Idade: 14},
			{Nome: "Eliana", Idade: 15},
			{Nome: "Gabriel", Idade: 14},
		},
	}

	// Imprimindo os nomes
	fmt.Printf("%s tem os alunos: %s, %s e %s\n",
		turma.NomeTurma,
		turma.Aluno[0].Nome,
		turma.Aluno[1].Nome,
		turma.Aluno[2].Nome,
	)
}

/*Turma com alunos

Crie uma struct Aluno com Nome e Idade.

Crie uma struct Turma com NomeTurma e uma slice de Aluno.

No main, crie uma turma com 3 alunos e imprima:
"Turma X tem os alunos: A, B, C".*/
