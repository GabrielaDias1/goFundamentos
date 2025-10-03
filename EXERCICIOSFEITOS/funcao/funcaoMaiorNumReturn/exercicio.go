package main

import "fmt"

func maior(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maior(10, 7))
}

/*Função que retorna o maior número

Crie uma função maior(a int, b int) int que retorna o maior dos dois números.*/
