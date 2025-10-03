package main

import "fmt"

func soma(s ...int) int {

	soma := 0

	for _, v := range s {
		soma += v

	}
	return soma
}

func main() {
	numero := []int{1, 2, 3, 4, 5}

	resultado := soma(numero...)

	fmt.Println("Soma =", resultado)
}

//Crie uma função que receba um slice e retorne a soma de todos os números.
