package main

import "fmt"

func ehPar(numero int) bool {
	return numero%2 == 0
}

func imprimirResultado(numero int) {
	if ehPar(numero) {
		fmt.Printf("O número %d é par\n", numero)
	} else {
		fmt.Printf("O número %d é ímpar\n", numero)
	}
}

func main() {
	imprimirResultado(4)
	imprimirResultado(5)
}

/*Função que chama outra função

Crie uma função imprimirResultado(n int) que usa a função ehPar para imprimir:

"O número X é par"

"O número X é ímpar"*/
