package main

import "fmt"

func main() {
	numeros := [7]int{1, 2, 3, 4, 5, 6, 7}

	soma := 0
	for i := 0; i <= 6; i++ {// ate indice 6
		soma += numeros[i]

	}
	fmt.Println("soma: ", soma)
}

/* Iteração e soma
Declare um array com 7 números inteiros quaisquer. Crie um laço para somar todos os números e depois imprima o resultado.*/
