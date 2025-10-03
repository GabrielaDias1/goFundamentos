package main

import "fmt"

func Contagem(n int, c ...int) int {

	conta := 0

	for _, v := range c {
		if v == n {
			conta++
		}
	}
	return conta

}

func main() {
	numero := []int{1, 2, 3, 2, 4, 2}

	resultado := Contagem(2, numero...)
	fmt.Println("O número 2 aparece", resultado, "vezes")

}

/*Crie uma função que receba um slice de números e um número extra.
A função deve contar quantas vezes esse número aparece no slice.*/
