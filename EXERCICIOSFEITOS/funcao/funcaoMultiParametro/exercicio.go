package main

import "fmt"

func main() {
	fmt.Println(somar(5, 4))
}

func somar(a int, b int) int {
	soma := a + b
	return soma
}

/*Função com múltiplos parâmetros

Crie uma função chamada somar(a int, b int) int que recebe dois números inteiros e retorna a soma deles.*/
