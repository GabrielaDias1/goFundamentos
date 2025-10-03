package main

import "fmt"

func somaAteN(n int) int {

	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	return soma

}

func main() {
	fmt.Println(somaAteN(4))
}

/*Função com um parâmetro

Crie uma função que recebe um número inteiro n e retorna a soma de todos os números de 1 até n.*/
