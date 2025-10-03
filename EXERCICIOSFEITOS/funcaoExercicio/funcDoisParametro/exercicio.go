package main

import "fmt"

func numeroMaior(a int, b int) int {
	if a > b {
		return a

	}
	return b
}

func main() {
	fmt.Println(numeroMaior(4, 5))
}

/*Função com dois parâmetros

Crie uma função que recebe dois números inteiros e retorna o maior deles.*/
