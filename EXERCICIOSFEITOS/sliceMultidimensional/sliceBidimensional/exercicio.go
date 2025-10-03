package main

import "fmt"

func main() {
	numeros := [][]int{
		{1, 2},
		{3, 4},
	}
	for i, linha := range numeros { // percorre as linhas
		for j, valor := range linha { // percorre os elementos da linha
			fmt.Printf("numeros[%d][%d] = %d\n", i, j, valor)
		}
	}
}

//Crie um slice bidimensional [][]int representando a matriz:
//Imprima cada número junto com sua posição [linha][coluna].
