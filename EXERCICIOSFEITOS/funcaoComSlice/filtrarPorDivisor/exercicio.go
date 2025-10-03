package main

import "fmt"

func numsDivisel(num []int, divisor int, verifica bool) []int {
	resultado := []int{}

	for _, v := range num {
		if verifica {
			if v%divisor == 0 {
				resultado = append(resultado, v)
			}
		} else {
			if v%divisor != 0 {
				resultado = append(resultado, v)
			}
		}
	}

	return resultado
}

func main() {

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	div := numsDivisel(numeros, 2, true)
	fmt.Println(div)
}

/*Crie uma função que receba:

um slice de números,

um divisor (int),

um booleano (true/false).

Se o booleano for true, retorne todos os números divisíveis pelo divisor.
Se for false, retorne os que não são divisíveis.*/
