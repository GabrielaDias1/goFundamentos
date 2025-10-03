package main

import "fmt"

func aprovacaoFinal(nota1 float64, nota2 float64, nota3 float64) string {
	media := (nota1 + nota2 + nota3) / 3

	if media >= 6 {
		return fmt.Sprintf("Média: %.2f - Aprovado", media)
	} else if media >= 4 {
		return fmt.Sprintf("Média: %.2f - Recuperação", media)
	} else {
		return fmt.Sprintf("Média: %.2f - Reprovado", media)
	}
}

func main() {
	fmt.Println(aprovacaoFinal(5, 5, 5))
	fmt.Println(aprovacaoFinal(8, 6, 7))
	fmt.Println(aprovacaoFinal(3, 4, 2))
}

/*Crie uma função que recebe três notas de um aluno e retorna:

A média

Se o aluno foi Aprovado, Recuperação ou Reprovado
(Regras: média >= 6 aprovado, entre 4 e 6 recuperação, menor que 4 reprovado).*/
