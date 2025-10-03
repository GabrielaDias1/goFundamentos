package main

import "fmt"

func menorNum(a int, b int, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	}
	return c
}

func main() {
	fmt.Println(menorNum(4, 5, 6))
}

/*Função com três parâmetros

Crie uma função que recebe três números inteiros e retorna o menor deles.*/
