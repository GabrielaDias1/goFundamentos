package main

import "fmt"

type aluno struct {
	Nome  string
	Nota1 float64
	Nota2 float64
	Nota3 float64
}

func (a aluno) Media() float64 {
	media := (a.Nota1 + a.Nota2 + a.Nota3) / 3
	return media
}

func (a aluno) Resultado() string {
	media := a.Media()
	if media >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	alunoGabriela := aluno{
		Nome:  "Gabriela",
		Nota1: 5.0,
		Nota2: 8.0,
		Nota3: 3.0,
	}
	fmt.Printf("Média de %s: %.2f\n", alunoGabriela.Nome, alunoGabriela.Media())
	fmt.Printf("Resultado: %s\n", alunoGabriela.Resultado())
}

/*Crie uma struct Aluno com nome e três notas (n1, n2, n3).
Implemente um método Media() que calcula e retorna a média do aluno.
Depois, faça um método Resultado() que retorna "Aprovado" se a média ≥ 6, caso contrário "Reprovado".*/
