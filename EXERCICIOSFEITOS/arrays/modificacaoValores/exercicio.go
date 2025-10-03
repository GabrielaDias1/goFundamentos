package main

import "fmt"

func main() {
	numeros := [5]int{2, 4, 6, 8, 10}

	numeros[4] = 12

	fmt.Printf("substituição: %v",numeros)
}

/*Modificação de valores
Crie um array com 5 números pares e substitua o último número por um número ímpar. Imprima o array modificado.*/
