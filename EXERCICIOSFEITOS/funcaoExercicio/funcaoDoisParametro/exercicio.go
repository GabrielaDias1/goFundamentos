package main

import "fmt"

func mediaNota(nota1 float64, nota2 float64) string {
	media := (nota1 + nota2) / 2

	if media >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {

	fmt.Println(mediaNota(4.0, 8.0))
}

/*Função com dois parâmetros

Crie uma função que recebe dois números (nota1, nota2) e retorna a média deles.
Depois, no main, verifique se o aluno foi aprovado (média >= 6).*/
