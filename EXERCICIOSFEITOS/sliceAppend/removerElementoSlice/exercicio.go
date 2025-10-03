package main

import "fmt"

func main() {
	numeros := []int{10, 20, 30, 40}

	fmt.Println(numeros)

	numeros = append(numeros[:2], numeros[3:]...)

	fmt.Println(numeros)
}

/*Removendo um elemento

Crie um slice []int{10, 20, 30, 40}.
Remova o valor 30 usando.
Mostre o slice final.*/
