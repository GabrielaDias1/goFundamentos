package main

import "fmt"

func somaSlice(numeros []int) int {
	soma := 0
	for i, _ := range numeros {
		soma += numeros[i]
		fmt.Println("soma: ", soma)
	}
	return soma
}

func main() {
	fmt.Println(somaSlice([]int{1, 2, 3, 4, 5, 6}))
}

/*Função que recebe slice

Crie uma função somaSlice(numeros []int) int que recebe um slice de inteiros e retorna a soma de todos os elementos.*/
