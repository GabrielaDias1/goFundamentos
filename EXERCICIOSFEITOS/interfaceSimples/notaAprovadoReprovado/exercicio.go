package main

import "fmt"

type Resultado interface {
	Aprovado(nota float64) bool
}

type Aluno struct{}

func (a Aluno) Aprovado(nota float64) bool {
	if nota >= 6.0 {
		return true
	} else {
		return false
	}
}

func genericaInterface(r Resultado, nota float64) {
	if r.Aprovado(nota) {
		fmt.Println("Aluno aprovado!")
	} else {
		fmt.Println("Aluno não aprovado")
	}
}

func main() {
	Aluno12345 := Aluno{}

	genericaInterface(Aluno12345, 10.0)
}

/*Nota Aprovado/Reprovado

Descrição:
Crie uma interface chamada Resultado com um método:

Aprovado(nota float64) bool

Crie uma struct Aluno que implemente esse método.

Retorne true se a nota for maior ou igual a 6.0.

Retorne false caso contrário.

No main, teste com uma nota por vez.*/
