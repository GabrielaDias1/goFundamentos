package main

import "fmt"

func numIntervalo(nums []int, valorMin int, valorMax int) []int {
	numeros := []int{}

	for _, v := range nums {
		if v >= valorMin && v <= valorMax {
			numeros = append(numeros, v)
		}
	}
	return numeros

}

func main() {
	resultado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numsIntervalo := numIntervalo(resultado, 3, 7)
	fmt.Println(numsIntervalo)
}

/*Crie uma função que receba:
um slice de números,
um valor mínimo,
um valor máximo.
E retorne apenas os números dentro desse intervalo.*/
