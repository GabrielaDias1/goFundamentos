package main

import "fmt"

func main() {

	s, m := operacoes(4, 5)
	fmt.Println("Soma:", s)
	fmt.Println("Multiplicação:", m)
}

func operacoes(a int, b int) (int, int) {
	soma := a + b
	multi := a * b
	return soma, multi
}

/*Função com múltiplos retornos

Crie uma função chamada operacoes(a int, b int) (int, int) que recebe dois números inteiros e retorna a soma e o produto desses números.*/
