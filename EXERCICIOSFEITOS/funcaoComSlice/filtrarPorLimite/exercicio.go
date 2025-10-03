package main

import "fmt"

func numMaior(limite int, n ...int) []int {
	resultado := []int{}

	for _, c := range n {
		if c > limite {
			resultado = append(resultado, c)
		}
	}

	return resultado
}

func main() {
	numeros := []int{1, 5, 3, 7, 2, 10}
	maiores := numMaior(4, numeros...)
	fmt.Println("Números maiores que 4:", maiores)
}

/*Crie uma função que receba um slice de números e um limite.
Retorne apenas os números maiores que o limite em um novo slice.*/
