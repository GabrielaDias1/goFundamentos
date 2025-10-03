package main

import "fmt"

type Aluno struct {
	Nome     string
	Nota     float64
	Aprovado bool
}

func main() {
	aluno1 := Aluno{
		Nome:     "Carol",
		Nota:     10,
		Aprovado: true,
	}
	aluno2 := Aluno{
		Nome:     "Eduarda",
		Nota:     5,
		Aprovado: false,
	}

	fmt.Printf("O aluno %s tirou %.1f e aprovado = %t\n", aluno1.Nome, aluno1.Nota, aluno1.Aprovado)
	fmt.Printf("O aluno %s tirou %.1f e aprovado = %t\n", aluno2.Nome, aluno2.Nota, aluno2.Aprovado)

}

/*Ficha de aluno
Crie uma struct Aluno com Nome (string), Nota (float64) e Aprovado (bool).
Crie um aluno e exiba a seguinte mensagem:
"O aluno X tirou Y e aprovado = Z".*/
