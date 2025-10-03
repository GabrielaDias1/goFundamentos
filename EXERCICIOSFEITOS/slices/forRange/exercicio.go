package main

import "fmt"

func main() {
	cores := []string{"branco", "azul", "amarelo", "verde"}

	for indice, valor := range cores {
		fmt.Println("Indice e cor: ", indice, valor)
	}
}

/*Percorrendo com for range
Crie um slice com 4 cores e use for range para imprimir o índice e o valor de cada elemento.*/
