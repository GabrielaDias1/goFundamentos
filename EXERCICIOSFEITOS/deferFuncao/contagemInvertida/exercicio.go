package main

import "fmt"

func imprimirInvertidoComDefer() {

	numeros := []int{1, 2, 3, 4, 5}

	for _, v := range numeros {
		defer fmt.Println(v)
	}
}

func main() {
	imprimirInvertidoComDefer()
}

/*Faça uma função que use um for para imprimir os números de 1 a 5, mas usando defer para inverter a ordem da saída.*/
