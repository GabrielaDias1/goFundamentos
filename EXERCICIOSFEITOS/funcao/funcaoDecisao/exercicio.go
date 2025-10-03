package main

import "fmt"

func situacaoAluno(nota float64) string {
	if nota >= 6.0 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(situacaoAluno(7.5))
	fmt.Println(situacaoAluno(5.0))
}

/*Função com decisão interna

Crie uma função situacaoAluno(nota float64) string que retorna:

"Aprovado" se a nota for maior ou igual a 6.0

"Reprovado" caso contrário.*/
